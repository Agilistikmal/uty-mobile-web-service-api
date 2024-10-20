package model

type User struct {
	Username string `json:"username,omitempty" gorm:"primaryKey"`
	FullName string `json:"full_name,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}
