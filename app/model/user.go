package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"username"`
	Password string `json:"password"`
}

func CreateUser(user *User) bool {
	if err := db.Create(user).Error; err != nil {
		return false
	}
	return true
}

func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}
