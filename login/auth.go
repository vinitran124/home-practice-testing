package login

const (
	ValidUsername = "vinitran"
	ValidPassword = "vinideptrai"
)

func Login(username, password string) bool {
	if username != ValidUsername {
		return false
	}

	if password != ValidPassword {
		return false
	}

	return true
}
