package forum

import (
	"database/sql"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func DeleteCookie(w http.ResponseWriter, r *http.Request) { // fonction supprimer cookie
	cookie := http.Cookie{
		Name:   "session",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)

	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	defer db.Close()
	uuid := getCookie(r)
	_, err = db.Exec("DELETE FROM COOKIE WHERE UUID=:0", uuid)
	checkErr(err)
}

