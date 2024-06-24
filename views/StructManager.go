package forum

import (
	_ "github.com/mattn/go-sqlite3"
)

const (
	path = "../forum.db"
)

type USERS struct {
	Id             int
	Username       string
	Password       string
	Email          string
	Permissions_Id int
	Biography      string
	Image          string
}

type POST struct {
	ID         int
	Post_Name  string
	Content    string
	Like_Nb    int
	Dislike_Nb int
	Topics_Id  int
	Author     USERS
	Date_Pb    string
	Tag_Id     int
}

type COMMENT struct {
	Id         int
	Content    string
	Like_Nb    int
	Dislike_Nb int
	Posts_Id   int
	Author     USERS
	Date_Pb    string
}

type TAGS struct {
	ID   int
	Tags string
}
