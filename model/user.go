package model

type User struct {
	ID       int        `json:"id" example:"1" format:"int" minimum:"1"`
	Name     string     `json:"name" example:"John Doe"`
	Email    string     `json:"email" example:"john.doe@example.com"`
	Password string     `json:"password" example:"********"`
	Role     string     `json:"role" enum:"admin,user"`
	Kenangan []Kenangan `json:"kenangan" gorm:"foreignKey:UserID"`
}
