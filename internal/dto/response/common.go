package response

type Response[T any] struct {
	Status  bool     `json:"status"`
	Error   []string `json:"error"`
	Message string   `json:"message"`
	Data    []T      `json:"data" omitempty:""`
}
