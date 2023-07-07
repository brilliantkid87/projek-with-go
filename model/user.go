package model

// User represents a user entity
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`

	// Kenangan represents the list of memories associated with the user
	// @gorm:"foreignKey:UserID"
	Kenangan []Kenangan `json:"kenangan" gorm:"foreignKey:UserID"`
}
