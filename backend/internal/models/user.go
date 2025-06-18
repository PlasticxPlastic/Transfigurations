package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // Password won't be included in JSON responses
}

func (u *User) ValidateLogin(password string) bool {
	// TODO: Implement actual password validation
	// For now, just return true if passwords match
	return u.Password == password
} 