package user

type CreateUser struct {
	Name           string `json:"name"`
	DocumentNumber string `json:"documentNumber"`
	DocumentType   int    `json:"documentType"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

type CreatedUser struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	DocumentNumber string `json:"documentNumber"`
	DocumentType   int    `json:"documentType"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}
