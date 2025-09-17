package waifuim

import "time"

// Response is the top-level structure of the waifu.im API response.
type Response struct {
	Images []Image `json:"images"`
}

// Image is the structure of an image object from the waifu.im API.
type Image struct {
	ImageID   int       `json:"image_id"`
	URL       string    `json:"url"`
	Source    string    `json:"source"`
	Tags      []Tag     `json:"tags"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Favorites int       `json:"favorites"`
	IsNSFW    bool      `json:"is_nsfw"`
	UploadedAt time.Time `json:"uploaded_at"`
}

// Tag is the structure of a tag object from the waifu.im API.
type Tag struct {
	Name string `json:"name"`
}
