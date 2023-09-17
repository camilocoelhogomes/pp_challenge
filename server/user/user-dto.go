package user

type CreateUser struct {
	Name           string `json:"name"`
	DocumentNumber string `json:"documentNumber"`
	DocumentType   int    `json:"documentType"`
	Email          string `json:"email"`
	Password       string `json:"password"`
}

type CreatedUser struct {
	Id             string `json:"id" gorm:"primaryKey;collum:id;type:uuid;default:uuid_generate_v4()"`
	internalId     int64  `gorm:"primaryKey;collum:internal_id"`
	Name           string `json:"name" gorm:"collum:name"`
	DocumentNumber string `json:"documentNumber" gorm:"collum:document_number"`
	DocumentType   int    `json:"documentType" gorm:"collum:document_type"`
	Email          string `json:"email" gorm:"collum:email"`
}

type Tabler interface {
	TableName() string
}

func (CreatedUser) TableName() string {
	return "user"
}

func (c *CreatedUser) getInternalId() int64 {
	return c.internalId
}
