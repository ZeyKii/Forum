package forum

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func GetDataUser(column, value string) USERS { // fonction récupérer les données d'un utilisateur
	var Users USERS
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("select id, username, email, image from users where " + column + " =:0")
	checkErr(err)

	res, err := stmt.Query(value)
	checkErr(err)

	defer res.Close()

	for res.Next() {
		err := res.Scan(&Users.Id, &Users.Username, &Users.Email, &Users.Image)
		checkErr(err)
	}
	return Users
}

func getAllUsers() []USERS {
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, username, email, image FROM users")
	checkErr(err)

	defer rows.Close()

	var Users []USERS
	for rows.Next() {
		var user USERS
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Image)
		checkErr(err)
		Users = append(Users, user)
	}
	return Users
}

func GetDataUserToken(r *http.Request) USERS { // fonction récupérer les données token d'un utilisateur
	var userInfo USERS
	if CheckCookie(r) {
		db, err := sql.Open("sqlite3", path)
		checkErr(err)
		defer db.Close()

		stmt, err := db.Prepare("select user_id from cookie where UUID " + " =:0")
		checkErr(err)

		res, err := stmt.Query(getCookie(r))
		checkErr(err)

		for res.Next() {
			err := res.Scan(&userInfo.Id)
			if err != nil {
				checkErr(err)
			}
		}
		stmt, err = db.Prepare("select username, email, password, biography, permissions_id, image from users where id " + " =:0")
		checkErr(err)

		res, err = stmt.Query(&userInfo.Id)
		checkErr(err)

		for res.Next() {
			err := res.Scan(&userInfo.Username, &userInfo.Email, &userInfo.Password, &userInfo.Biography, &userInfo.Permissions_Id, &userInfo.Image)
			if err != nil {
				checkErr(err)
			}
		}
		return userInfo
	}
	return userInfo
}

func GetPosts() []POST { // fonction récupérer les posts dans la base de donnée
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM POSTS")
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)

	defer rows.Close()
	var post POST
	var Posts []POST
	for rows.Next() {
		var author_id int
		err := rows.Scan(&post.ID, &post.Content, &post.Post_Name, &post.Date_Pb, &author_id, &post.Tag_Id)
		checkErr(err)
		statement, err := db.Prepare("SELECT id, username, image FROM users where id = :0")
		checkErr(err)

		row, err := statement.Query(author_id)
		checkErr(err)

		defer row.Close()
		var user USERS
		for row.Next() {
			err = row.Scan(&user.Id, &user.Username, &user.Image)
			checkErr(err)
		}
		post.Author = user
		Posts = append(Posts, post)
	}
	return Posts
}

func GetPostId(column, value string) POST { // fonction récupérer les postsID dans la base de donnée
	var post POST
	db, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM POSTS WHERE " + column + " = " + value)
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var author_id int
		err := rows.Scan(&post.ID, &post.Content, &post.Post_Name, &post.Date_Pb, &author_id, &post.Tag_Id)
		checkErr(err)
		statement, err := db.Prepare("SELECT id, username, image FROM users where id = :0")
		checkErr(err)

		row, err := statement.Query(author_id)
		checkErr(err)

		defer row.Close()
		var user USERS
		for row.Next() {
			err = row.Scan(&user.Id, &user.Username, &user.Image)
			checkErr(err)
		}
		post.Author = user
	}
	return post
}

func GetPostUser(UserID string) []POST { // fonction récupérer les postsUser dans la base de donnée
	db, err := sql.Open("sqlite3", "../forum.db")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM POSTS WHERE User_ID = " + UserID)
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	defer rows.Close()
	var post POST
	var PostUser []POST
	for rows.Next() {
		var author_id int
		err := rows.Scan(&post.ID, &post.Content, &post.Post_Name, &post.Date_Pb, &author_id, &post.Tag_Id)
		checkErr(err)
		statement, err := db.Prepare("SELECT id, username, image FROM users where id = :0")
		checkErr(err)

		row, err := statement.Query(author_id)
		checkErr(err)

		defer row.Close()
		var user USERS
		for row.Next() {
			err = row.Scan(&user.Id, &user.Username, &user.Image)
			checkErr(err)
		}
		post.Author = user
		PostUser = append(PostUser, post)
	}
	return PostUser
}

func GetComment(column, value string) []COMMENT { // fonction récupérer les commentaires dans la base de donnée
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("select * from comments where " + column + " =:0")
	checkErr(err)

	rows, err := stmt.Query(value)
	checkErr(err)

	defer rows.Close()
	var comment COMMENT
	var comments []COMMENT
	var author_id int
	for rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Content, &comment.Posts_Id, &author_id, &comment.Date_Pb)
		if err != nil {
			checkErr(err)
		}
		statement, err := db.Prepare("SELECT id, username, image FROM users where id = :0")
		checkErr(err)

		row, err := statement.Query(author_id)
		checkErr(err)

		defer row.Close()
		var user USERS
		for row.Next() {
			err = row.Scan(&user.Id, &user.Username, &user.Image)
			checkErr(err)
		}
		comment.Author = user
		comments = append(comments, comment)
	}
	return comments
}

func GetTag(id int) string { // fonction récupérer les tags dans la base de donnée
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	var tag string
	err = db.QueryRow("SELECT TAG FROM TAGS WHERE ID=:0", id).Scan(&tag)
	checkErr(err)
	return tag
}

func getCookie(r *http.Request) string { // fonction récupérer les cookies dans la base de donnée
	cookie, err := r.Cookie("session")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func GetLikePost(id int) {
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	var like int
	err = db.QueryRow("SELECT COUNT(*) FROM LIKES_POST WHERE POST_ID=:0", id).Scan(&like)
	checkErr(err)
	var post POST
	post.Like_Nb = like
}

func GetDislikePost(id int) {
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	var dislike int
	err = db.QueryRow("SELECT COUNT(*) FROM DISLIKES_POST WHERE POST_ID=:0", id).Scan(&dislike)
	checkErr(err)
	var post POST
	post.Dislike_Nb = dislike
}
