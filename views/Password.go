package forum

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string { // fonction Hasher un password
	pw := []byte(password) // On utilise des tableaux de bytes pour Bcrypt
	result, _ := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	return string(result)
}

func ComparePassword(hashPassword string, password string) error { // fonction compare un hash de MDP au string d'un MDP user
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	if err != nil {
		checkErr(err)
	}
	return err
}
