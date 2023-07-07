package authdto

type LoginResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Role  string `gorm:"type: varchar(255)" json:"role"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Role  string `gorm:"type: varchar(255)" json:"role"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
