package nekosapi

import "time"

// Response is the top-level structure of the nekosapi.com API response.
type Response struct {
	Data []Image `json:"data"`
}

// Image is the structure of an image object from the nekosapi.com API.
type Image struct {
	ID        string    `json:"id"`
	Rating    string    `json:"rating"`
	Tags      []Tag     `json:"tags"`
	Artist    *Artist   `json:"artist"`
	Source    Source    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	FileURL   string    `json:"file_url"`
}

// Tag is the structure of a tag object.
type Tag struct {
	Name string `json:"name"`
}

// Artist is the structure of an artist object.
type Artist struct {
	Name string `json:"name"`
}

// Source is the structure of a source object.
type Source struct {
	URL string `json:"url"`
}
