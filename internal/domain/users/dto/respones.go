package dto

// User Response
type Response struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role  string `json:"role"`
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"created_at"`
}

