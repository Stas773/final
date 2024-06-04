package entities

type Result struct {
	Status bool      `json:"status"`
	Data   ResultSet `json:"data"`
	Error  string    `json:"error"`
}
