package waifupics

// request is the request body for the waifu.pics API.
type request struct {
	Exclude []string `json:"exclude"`
}

// response is the response body from the waifu.pics API.
type response struct {
	Files []string `json:"files"`
}
