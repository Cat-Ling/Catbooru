package picre

// Image is the structure of an image object from the pic.re API.
type Image struct {
	ID      int      `json:"_id"`
	FileURL string   `json:"file_url"`
	Source  string   `json:"source"`
	Tags    []string `json:"tags"`
	Width   int      `json:"width"`
	Height  int      `json:"height"`
}
