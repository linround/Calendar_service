package utils

func BcryptCheck(password, hash string) bool {
	return password == hash
}
