package model

type Auth struct {
	ID       int `gorm:"primary_key"`
	Email    string
	Password string
}

func CreateAuth(auth *Auth) error {
	err := db.Create(auth).Error
	if err != nil {
		return err
	}
	return nil
}

func FetchByEmailAndPassword(a *Auth) Auth {
	var auth Auth
	db.Where(&Auth{Email: a.Email, Password: a.Password}).First(&auth)
	return auth
}
