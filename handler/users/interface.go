package users

type NewUsersRepository interface {
	Registration(payload Payload) error
	Update(payload Payload) error
	Profile(id string) (Payload, error)
	Login(phone_number string) (Payload, error)
}
