package handlers

const (
	InternalError      = "internal server error"
	FormError          = "incorrect form"
	PasswordTooShort   = "password too short, at least 8 letters"
	NoError            = "ok"
	NoEmail            = "no email field"
	NoUsername         = "no username field"
	NoPassword         = "no password field"
	NoRepeatedPassword = "no repeated password field"
)

type Error struct {
	Message string `json:"message"`
}
