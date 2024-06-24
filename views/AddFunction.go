package forum

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func AddUser(username, password, email string) { // fonction ajouter un utilisateur
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (username, password, email) values(:0, :1, :2)", username, HashPassword(password), email)
	checkErr(err)

}

func AddUserPic(column, username string) { // fonction ajouter une image d'utilisateur
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	ImagePath := "/assets/users/upload-" + username + ".png"

	_, err = db.Exec("UPDATE users SET IMAGE = ? WHERE username = ?", ImagePath, username)
	fmt.Println(err)
	checkErr(err)
}

func AddCookie(Id int, uuid string) { // fonction ajouter un cookie à un utilisateur
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	_, err = db.Exec("INSERT INTO cookie (UUID, User_id) values(:0, :1)", uuid, Id)
	checkErr(err)

}

func AddPost(name string, content string, r *http.Request, id_tag int) { // fonction ajouter un post
	local := time.Now().Local().UTC().Add(1 * time.Hour)
	parsedate := local.Format("02/01/2006 à 15h04")
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	uuid := getCookie(r)
	switch r.Method {
	case "POST":
		value := r.FormValue("tags")
		DbHTML.Tag.Tags = value
		switch DbHTML.Tag.Tags {
		case "Général":
			id_tag = 1
		case "Chasse_au_trésor":
			id_tag = 2
		case "Théorie":
			id_tag = 3
		}
		var User_id string
		err = db.QueryRow("SELECT USER_ID FROM COOKIE WHERE UUID=:0", uuid).Scan(&User_id)
		checkErr(err)
		_, err = db.Exec("INSERT INTO POSTS (NAME, CONTENT, DATE_PB, USER_ID, TAG_ID) values(:0, :1, :2, :3, :4)", name, content, parsedate, User_id, id_tag)
		checkErr(err)
	}
}

func AddComment(Id int, content string, w http.ResponseWriter, r *http.Request) { // fonction ajouter un commentaire à un post
	if !CheckCookie(r) {
		http.Redirect(w, r, "/authentification", http.StatusSeeOther)
		return
	}
	local := time.Now().UTC().Add(1 * time.Hour)
	parsedate := local.Format("02/01/2006 à 15h04")
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	uuid := getCookie(r)
	var User_id string
	err = db.QueryRow("SELECT USER_ID FROM COOKIE WHERE UUID=:0", uuid).Scan(&User_id)
	checkErr(err)
	_, err = db.Exec("INSERT INTO COMMENTS (CONTENT, DATE_PB, USER_ID, POSTS_ID) values(:0, :1, :2, :3)", content, parsedate, User_id, Id)
	checkErr(err)
}

func AddBio(bio string, user_id int, r *http.Request) { // fonction ajouter une bio d'utilisateur
	Id := strconv.Itoa(user_id)
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	_, err = db.Exec("UPDATE USERS SET BIOGRAPHY = :0 WHERE ID = :1", bio, Id)
	checkErr(err)
}
