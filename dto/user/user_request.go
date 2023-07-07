package userdto

type CreateUserRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Kenangan []model.KenanganUser `json:"kenangans"`
}
