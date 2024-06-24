package forum

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func PasswordCheck(username, password string) bool { // fonction vérifier bon mot de passe
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("select password from users where username=:0")
	checkErr(err)

	res, err := stmt.Query(username)
	checkErr(err)

	defer res.Close()
	var userPass string
	for res.Next() {
		err := res.Scan(&userPass)
		if err != nil {
			checkErr(err)
		}
	}
	fmt.Println("userpass =>", userPass)
	fmt.Println("password =>", password)
	fmt.Println("compare : ", ComparePassword(userPass, password) == nil)
	return ComparePassword(userPass, password) == nil
}

func RemoveUser(username string) { // fonction enlever un utilisateur
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE username=:0", username)
	checkErr(err)
}

func EmailCheck(Email string) bool {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM USERS WHERE EMAIL = :0", Email).Scan(&count)
	if err != nil {
		checkErr(err)
	}

	return count > 0
}

func UsernameCheck(username string) bool { // fonction vérifier bon nom d'utilisater
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM USERS WHERE USERNAME = :0", username).Scan(&count)
	if err != nil {
		checkErr(err)
	}

	return count > 0
}

func checkErr(err error) { // fonction gestion d'erreur
	if err != nil {
		checkErr(err)
	}
}

func CheckCookie(r *http.Request) bool { // fonction vérifier si l'utilisateur à un cookie
	_, err := r.Cookie("session")
	if err != nil {
		return false
	}
	return true
}
