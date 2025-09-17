package server

import (
	"booru-server/internal/config"
	"booru-server/pkg/booru"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/time/rate"
)

// Server is the main application server.
type Server struct {
	router    *chi.Mux
	clients   []booru.BooruClient
	authCfg   config.AuthConfig
	limiter   *rate.Limiter
	rateLimitEnabled bool
}

// New creates a new server with the given clients and configs.
func New(clients []booru.BooruClient, authCfg config.AuthConfig, rateLimitCfg config.RateLimitConfig) *Server {
	s := &Server{
		router:    chi.NewRouter(),
		clients:   clients,
		authCfg:   authCfg,
		rateLimitEnabled: rateLimitCfg.Enabled,
	}

	if s.rateLimitEnabled {
		s.limiter = rate.NewLimiter(rate.Limit(rateLimitCfg.RequestsPerSecond), rateLimitCfg.Burst)
	}

	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.With(s.authMiddleware, s.rateLimitMiddleware).Get("/api/search", s.handleSearch())

	return s
}

// Start starts the HTTP server on the given address.
func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.authCfg.Enabled {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := tokenParts[1]

		for _, validToken := range s.authCfg.Tokens {
			if token == validToken {
				next.ServeHTTP(w, r)
				return
			}
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (s *Server) rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !s.rateLimitEnabled {
			next.ServeHTTP(w, r)
			return
		}

		if !s.limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// @Summary Search for images
// @Description Searches for images across multiple booru sources.
// @Tags images
// @Accept json
// @Produce json
// @Param tags query string false "A comma-separated list of tags to search for."
// @Param nsfw query boolean false "Whether to include NSFW content."
// @Param limit query int false "The maximum number of results to return from each provider."
// @Param width query int false "A minimum width for the images."
// @Param height query int false "A minimum height for the images."
// @Param orderBy query string false "The order in which to sort the results."
// @Security Bearer
// @Success 200 {array} booru.Image
// @Failure 401 {string} string "Unauthorized"
// @Failure 429 {string} string "Too Many Requests"
// @Router /api/search [get]
func (s *Server) handleSearch() http.HandlerFunc {
	type result struct {
		images []booru.Image
		err    error
	}

	return func(w http.ResponseWriter, r *http.Request) {
		params := parseSearchParams(r)

		var wg sync.WaitGroup
		resultsChan := make(chan result, len(s.clients))

		for _, client := range s.clients {
			wg.Add(1)
			go func(c booru.BooruClient) {
				defer wg.Done()
				images, err := c.Search(r.Context(), params)
				if err != nil {
					err = fmt.Errorf("client %s: %w", c.Name(), err)
				}
				resultsChan <- result{images: images, err: err}
			}(client)
		}

		wg.Wait()
		close(resultsChan)

		var allImages []booru.Image
		for res := range resultsChan {
			if res.err != nil {
				log.Printf("Error from client: %v", res.err)
			} else {
				allImages = append(allImages, res.images...)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(allImages); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func parseSearchParams(r *http.Request) booru.SearchParams {
	q := r.URL.Query()
	var params booru.SearchParams

	if tags := q.Get("tags"); tags != "" {
		params.Tags = strings.Split(tags, ",")
	}
	if nsfw := q.Get("nsfw"); nsfw != "" {
		if b, err := strconv.ParseBool(nsfw); err == nil {
			params.NSFW = &b
		}
	}
	if limit := q.Get("limit"); limit != "" {
		if i, err := strconv.Atoi(limit); err == nil {
			params.Limit = i
		}
	}
	if width := q.Get("width"); width != "" {
		if i, err := strconv.Atoi(width); err == nil {
			params.Width = &i
		}
	}
	if height := q.Get("height"); height != "" {
		if i, err := strconv.Atoi(height); err == nil {
			params.Height = &i
		}
	}
	if orderBy := q.Get("orderBy"); orderBy != "" {
		params.OrderBy = orderBy
	}

	return params
}
