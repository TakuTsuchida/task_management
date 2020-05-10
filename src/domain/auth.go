package domain

type (
	Signup struct {
		Email    string
		Password string
	}
)

func (s Signup) ValidateSignup() bool {
	return s.Email != ""&s.Password != ""
}
