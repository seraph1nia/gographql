package auth

type Signuprequest struct {
	Username string `json:*username`
	Password string `json:*password`
	Email    string `json:*email`
}

func Signup(r Signuprequest) string {

	return Signuprequest.Email
}
