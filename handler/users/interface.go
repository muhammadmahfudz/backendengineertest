package users

type NewUsersRepository interface {
	Registration(payload Payload) error
	Update(payload Payload) error
	Profile(id string) *Payload
}
