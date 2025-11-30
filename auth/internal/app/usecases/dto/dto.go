package dto

type Credentials struct {
	Email    string
	Password string
}

type Register struct {
	Credentials
}

type Login struct {
	Credentials
}

type Refresh struct {
	Token string
}
