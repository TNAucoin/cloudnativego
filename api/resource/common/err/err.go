package err

type Error struct {
	Error error `json:"error"`
}

type Errors struct {
	Error []string `json:"errors"`
}
