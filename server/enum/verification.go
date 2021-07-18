package enum

type VerificationType int

const (
	BasicAuthSignup VerificationType = iota
	UpdateEmail
)

func (d VerificationType) String() string {
	return [...]string{
		"basic_auth_signup",
		"update_email",
	}[d]
}