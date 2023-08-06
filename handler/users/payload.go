package users

type Payload struct {
	ID          uint   `json:"id,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	FullName    string `json:"fullname,omitempty"`
	Password    string `json:"password,omitempty"`
}

type LoginResponse struct {
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
	Err     interface{} `json:"err,omitempty"`
}
