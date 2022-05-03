package apilayer

// BaseAPILayer base struct for checking the success of request.
type BaseAPILayer struct {
	Success bool  `json:"success"`
	Error   Error `json:"error"`
}

// Error structure response error.
type Error struct {
	Code int    `json:"code"`
	Info string `json:"info"`
}
