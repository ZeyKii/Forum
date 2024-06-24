package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func FilterPostbyLike() []POST { // fonction filtrer les posts par le nombre de like
	var Posts []POST
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM LIKES_POST ORDER BY SELECT COUNT(*) FROM LIKES_POST WHERE POST_ID = :0 DESC")
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)
	defer rows.Close()
	var post POST
	for rows.Next() {
		var author_id int
		err := rows.Scan(&post.ID, &post.Content, &post.Post_Name, &post.Date_Pb, &author_id, &post.Tag_Id)
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
		post.Author = user
		Posts = append(Posts, post)
	}
	return Posts
}

func FilterPostbyDate() []POST { // fonction filtrer les posts par la date
	var Posts []POST
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM POSTS ORDER BY DATE_PB DESC")
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)
	defer rows.Close()
	var post POST
	for rows.Next() {
		var author_id int
		err := rows.Scan(&post.ID, &post.Content, &post.Post_Name, &post.Date_Pb, &author_id, &post.Tag_Id)
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
		post.Author = user
		Posts = append(Posts, post)
	}
	return Posts
}

func FilterPostbyUsername(username_id, tag_id string) []POST { // fonction filtrer les posts par leurs noms
	var author_id USERS
	var Posts []POST
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM POSTS WHERE USER_ID = " + username_id)
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)
	defer rows.Close()
	var post POST
	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Content, &post.Post_Name, &post.Like_Nb, &post.Dislike_Nb, &post.Date_Pb, &author_id, &post.Tag_Id)
		if err != nil {
			checkErr(err)
		}
		statement, err := db.Prepare("SELECT * FROM POSTS WHERE TAG_ID = " + tag_id)
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

func FilterPostbyTag(tag_id string) []POST { // fonction filtrer les posts par le tag
	var Posts []POST
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM POSTS WHERE TAG_ID = " + tag_id)
	checkErr(err)

	rows, err := stmt.Query()
	checkErr(err)
	defer rows.Close()
	var post POST
	for rows.Next() {
		var author_id int
		err := rows.Scan(&post.ID, &post.Content, &post.Post_Name, &post.Date_Pb, &author_id, &post.Tag_Id)
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
		post.Author = user
		Posts = append(Posts, post)
	}
	return Posts
}
