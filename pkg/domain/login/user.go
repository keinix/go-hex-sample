package login

type User struct {
	Id           int64  `gorm:"PRIMARY_KEY" json:"id"`
	Username     string `json:"username"`
	PasswordHash string
}
