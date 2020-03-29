package model

type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"userName"`
	Password string `json:"password"`
}

func CreateUser(user *User) bool {
	err := db.Create(user).Error
	if err != nil {
		return false
	}
	return true
}

func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}
