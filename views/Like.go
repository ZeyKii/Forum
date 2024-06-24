package forum

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func LikePost(post_id int, user_id int) { // fonction aimer un post
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	// Check if user has already liked the post
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM LIKES_POST WHERE POST_ID = :0 AND USER_ID = :1", post_id, user_id).Scan(&count)
	checkErr(err)
	if count == 0 {
		_, err = db.Exec("INSERT INTO LIKES_POST (POST_ID, USER_ID) VALUES (:0, :1)", post_id, user_id)
		checkErr(err)
	} else {
		_, err = db.Exec("DELETE FROM LIKES_POST WHERE POST_ID = :0 AND USER_ID = :1", post_id, user_id)
		checkErr(err)
	}
}

func DislikePost(post_id int, user_id int) { // fonction ne pas aimer un post
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	// Check if user has already disliked the post
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM DISLIKES_POST WHERE POST_ID = :0 AND USER_ID = :1", post_id, user_id).Scan(&count)
	checkErr(err)
	if count == 0 {
		_, err = db.Exec("INSERT INTO DISLIKES_POST (POST_ID, USER_ID) VALUES (:0, :1)", post_id, user_id)
		checkErr(err)
	} else {
		_, err = db.Exec("DELETE FROM DISLIKES_POST WHERE POST_ID = :0 AND USER_ID = :1", post_id, user_id)
		checkErr(err)
	}
}

func LikeComment(comment_id int, user_id int) { // fonction aimer un commentaire
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	// Check if user has already liked the comment
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM LIKES_COMMENT WHERE COMMENT_ID = :0 AND USER_ID = :1", comment_id, user_id).Scan(&count)
	checkErr(err)
	if count == 0 {
		_, err = db.Exec("INSERT INTO LIKES_COMMENT (COMMENT_ID, USER_ID) VALUES (:0, :1)", comment_id, user_id)
		checkErr(err)
	} else {
		_, err = db.Exec("DELETE FROM LIKES_COMMENT WHERE COMMENT_ID = :0 AND USER_ID = :1", comment_id, user_id)
		checkErr(err)
	}
}

func DislikeComment(comment_id int, user_id int) { // fonction ne pas aimer un commentaire
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	// Check if user has already disliked the comment
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM DISLIKES_COMMENT WHERE COMMENT_ID = :0 AND USER_ID = :1", comment_id, user_id).Scan(&count)
	checkErr(err)
	if count == 0 {

		_, err = db.Exec("INSERT INTO DISLIKES_COMMENT (COMMENT_ID, USER_ID) VALUES (:0, :1)", comment_id, user_id)
		checkErr(err)
	} else {
		_, err = db.Exec("DELETE FROM DISLIKES_COMMENT WHERE COMMENT_ID = :0 AND USER_ID = :1", comment_id, user_id)
		checkErr(err)
	}
}