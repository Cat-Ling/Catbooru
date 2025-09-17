package booru

import "context"

// BooruModule is the interface that all booru API modules must implement.
type BooruModule interface {
	// Name returns the name of the booru provider (e.g., "waifu.im").
	Name() string
	// Search queries the booru API with the given parameters.
	Search(ctx context.Context, params SearchParams) ([]Image, error)
}
