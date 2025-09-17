package purrbotsite

// Response is the structure of the purrbot.site API response.
type Response struct {
	Link  string `json:"link"`
	Error bool   `json:"error"`
}
