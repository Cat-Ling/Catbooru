package server

import (
	"booru-server/pkg/booru"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server is the main application server.
type Server struct {
	router  *chi.Mux
	modules []booru.BooruModule
	version string
}

// New creates a new server with the given modules.
func New(modules []booru.BooruModule, version string) *Server {
	s := &Server{
		router:  chi.NewRouter(),
		modules: modules,
		version: version,
	}

	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))

	s.router.Get("/api/search", s.handleSearch())
	s.router.Get("/version", s.handleVersion())

	return s
}

// handleVersion returns the server's version.
func (s *Server) handleVersion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(s.version))
	}
}

// Start starts the HTTP server on the given address.
func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
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
// @Success 200 {array} booru.Image
// @Router /api/search [get]
func (s *Server) handleSearch() http.HandlerFunc {
	type result struct {
		moduleName string
		images     []booru.Image
		err        error
	}

	return func(w http.ResponseWriter, r *http.Request) {
		params := parseSearchParams(r)

		var wg sync.WaitGroup
		resultsChan := make(chan result, len(s.modules))

		for _, module := range s.modules {
			wg.Add(1)
			go func(m booru.BooruModule) {
				defer wg.Done()
				images, err := m.Search(r.Context(), params)
				resultsChan <- result{moduleName: m.Name(), images: images, err: err}
			}(module)
		}

		wg.Wait()
		close(resultsChan)

		var allImages []booru.Image
		for res := range resultsChan {
			if res.err != nil {
				log.Printf("Error from module '%s': %v", res.moduleName, res.err)
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
