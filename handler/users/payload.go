package users

type Payload struct {
	ID          uint   `json:"id"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	FullName    string `json:"fullname" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
