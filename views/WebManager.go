package forum

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/google/uuid"
)

type HTML struct {
	Users     USERS
	HomeUser  []USERS
	Post      POST
	Posts     []POST
	PostsUser []POST
	Comment   COMMENT
	Comments  []COMMENT
	Tag       TAGS
	Connexion string
	ID_Url    int
}

var DbHTML HTML = HTML{}

func StartServer() {
	fs := http.FileServer(http.Dir("../static/dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))
	img := http.FileServer(http.Dir("../assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", img))

	http.HandleFunc("/redirect", gotoAuth)
	http.HandleFunc("/authentification", homeHandler)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)

	http.HandleFunc("/", Forum)

	http.HandleFunc("/topic", Topic)
	http.HandleFunc("/posting", doPost)

	http.HandleFunc("/post", TmpPost)

	// http.HandleFunc("/likepost", UsingLikePost)
	// http.HandleFunc("/likecomment", UsingLikeComment)

	http.HandleFunc("/profile", Profile)
	http.HandleFunc("/upload", UploadFile)
	http.ListenAndServe(":8082", nil)
}

// -------------------------------------------- Chargement des différentes pages

func homeHandler(w http.ResponseWriter, r *http.Request) {
	userInfo := GetDataUserToken(r)
	var connexion string
	if !CheckCookie(r) {
		userInfo.Image = "/assets/ProfileIcon.png"
		connexion = "Connexion"
	} else {
		connexion = "Déconnexion"
	}
	if r.URL.Path != "/authentification" {
		http.Redirect(w, r, "https://http.cat/404", http.StatusSeeOther)
		return
	} else {
		cookie := http.Cookie{
			Name:   "session",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)
		files := []string{"../static/login.html", "../static/navbar.html"}
		tmp, err := template.ParseFiles(files...)
		checkErr(err)
		data := HTML{Users: userInfo, Connexion: connexion}
		err = tmp.Execute(w, data)
		if err != nil {
			checkErr(err)
		}
	}
}

func Forum(w http.ResponseWriter, r *http.Request) {
	userInfo := GetDataUserToken(r)
	var connexion string
	var AllUser []USERS
	AllUser = getAllUsers()
	if !CheckCookie(r) {
		userInfo.Image = "/assets/ProfileIcon.png"
		connexion = "Connexion"
	} else {
		connexion = "Déconnexion"
	}
	files := []string{"../static/forum.html", "../static/navbar.html"}
	tmpl := template.Must(template.ParseFiles(files...))
	data := HTML{Users: userInfo, Connexion: connexion, HomeUser: AllUser}
	err := tmpl.Execute(w, data)
	if err != nil {
		checkErr(err)
	}
}

func Topic(w http.ResponseWriter, r *http.Request) {
	userInfo := GetDataUserToken(r)
	var connexion string
	if !CheckCookie(r) {
		userInfo.Image = "/assets/ProfileIcon.png"
		connexion = "Connexion"
	} else {
		connexion = "Déconnexion"
	}
	var Posts = GetPosts()
	if r.Method == "POST" {
		switch r.FormValue("filter") {
		case "0":
			var Posts = GetPosts()
			files := []string{"../static/topics.html", "../static/navbar.html"}
			tmpl := template.Must(template.ParseFiles(files...))
			data := HTML{Users: userInfo, Connexion: connexion, Posts: Posts}
			tmpl.Execute(w, data)
		case "1":
			var Posts = FilterPostbyTag("1")
			files := []string{"../static/topics.html", "../static/navbar.html"}
			tmpl := template.Must(template.ParseFiles(files...))
			data := HTML{Users: userInfo, Connexion: connexion, Posts: Posts}
			tmpl.Execute(w, data)
		case "2":
			var Posts = FilterPostbyTag("2")
			files := []string{"../static/topics.html", "../static/navbar.html"}
			tmpl := template.Must(template.ParseFiles(files...))
			data := HTML{Users: userInfo, Connexion: connexion, Posts: Posts}
			tmpl.Execute(w, data)
		case "3":
			var Posts = FilterPostbyTag("3")
			files := []string{"../static/topics.html", "../static/navbar.html"}
			tmpl := template.Must(template.ParseFiles(files...))
			data := HTML{Users: userInfo, Connexion: connexion, Posts: Posts}
			tmpl.Execute(w, data)
		}
	} else {
		files := []string{"../static/topics.html", "../static/navbar.html"}
		tmpl := template.Must(template.ParseFiles(files...))
		data := HTML{Users: userInfo, Connexion: connexion, Posts: Posts}
		tmpl.Execute(w, data)
	}
}

func TmpPost(w http.ResponseWriter, r *http.Request) {
	userInfo := GetDataUserToken(r)
	IdPost := r.FormValue("id")
	id, _ := strconv.Atoi(IdPost)
	GetLikePost(id)
	GetDislikePost(id)
	var connexion string
	if !CheckCookie(r) {
		userInfo.Image = "/assets/ProfileIcon.png"
		connexion = "Connexion"
	} else {
		connexion = "Déconnexion"
	}
	if r.Method == "POST" {
		if r.FormValue("button") == "comment" {
			AddComment(id, r.FormValue("comment-content"), w, r)
		}
		if r.FormValue("button") == "like" {
			LikePost(id, GetDataUserToken(r).Id)
		} else if r.FormValue("button") == "dislike" {
			DislikePost(id, GetDataUserToken(r).Id)
		}
	}
	files := []string{"../static/post.html", "../static/navbar.html"}
	Post := GetPostId("id", IdPost)
	Comments := GetComment("Posts_Id", IdPost)
	tmpl := template.Must(template.ParseFiles(files...))
	data := HTML{Users: userInfo, Connexion: connexion, Post: Post, Comments: Comments}
	tmpl.Execute(w, data)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	var connexion string
	userInfo := GetDataUserToken(r)
	files := []string{"../static/profile.html", "../static/navbar.html"}
	if !CheckCookie(r) {
		userInfo.Image = "/assets/ProfileIcon.png"
		connexion = "Connexion"
	} else {
		connexion = "Déconnexion"
	}
	if !CheckCookie(r) {
		http.Redirect(w, r, "/authentification", http.StatusSeeOther)
	}
	if r.Method == "POST" {
		AddBio(r.FormValue("bio-content"), userInfo.Id, r)
	}
	data := HTML{Users: userInfo, Connexion: connexion}
	tmpl := template.Must(template.ParseFiles(files...))
	tmpl.Execute(w, data)
}

// -------------------------------------------- Fonction Ajout d'un cookie a l'utilisateur

func SetCookie(w http.ResponseWriter) http.Cookie {
	id := uuid.New()
	cookie := http.Cookie{
		Name:       "session",
		Value:      id.String(),
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     99999999999,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   []string{},
	}
	http.SetCookie(w, &cookie)
	return cookie
}

// -------------------------------------------- Fonction de longin/register

func Login(w http.ResponseWriter, r *http.Request) {
	var userInfo USERS
	cookie := SetCookie(w)
	// if r.Method != "POST" {
	// 	http.Redirect(w, r, "https://http.cat/405", http.StatusSeeOther)
	// 	return
	// }
	if r.Method == "POST" {
		userInfo.Username = r.FormValue("Username")
		userInfo.Password = r.FormValue("Password")
		GetDataUser("username", userInfo.Username)
		userInfo.Image = GetDataUser("username", userInfo.Username).Image
		if !UsernameCheck(userInfo.Username) || !PasswordCheck(userInfo.Username, userInfo.Password) {
			fmt.Println("Not OK")
			DeleteCookie(w, r)
			http.Redirect(w, r, "/login", 303)
			return
		}
		AddCookie(GetDataUser("username", userInfo.Username).Id, cookie.Value)
		http.Redirect(w, r, "/", 303)
		return
	} else {
		if CheckCookie(r) {
			userInfo = GetDataUserToken(r)
		}
	}
	data := HTML{Users: userInfo}
	files := []string{"../static/login.html", "../static/navbar.html"}
	tmp, err := template.ParseFiles(files...)
	checkErr(err)
	err = tmp.Execute(w, data)
	checkErr(err)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Redirect(w, r, "https://http.cat/405", http.StatusSeeOther)
	// 	return
	// }
	var userInfo USERS
	if r.Method == "POST" {
		userInfo.Email = r.FormValue("Email")
		userInfo.Username = r.FormValue("Username")
		userInfo.Password = r.FormValue("Password")
		if !EmailCheck(userInfo.Email) && !UsernameCheck(userInfo.Username) {
			AddUser(userInfo.Username, userInfo.Password, userInfo.Email)
			cookie := SetCookie(w)
			AddCookie(GetDataUser("username", userInfo.Username).Id, cookie.Value)
			userInfo.Image = GetDataUser("username", userInfo.Username).Image
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/authentification", http.StatusSeeOther)
	} else {
		if CheckCookie(r) {
			userInfo = GetDataUserToken(r)
		}

	}
	data := HTML{Users: userInfo}
	files := []string{"../static/login.html", "../static/navbar.html"}
	tmp, err := template.ParseFiles(files...)
	checkErr(err)
	err = tmp.Execute(w, data)
	checkErr(err)
}

// -------------------------------------------- Fonction de Redirection pour l'authentification

func gotoAuth(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w, r)
	http.Redirect(w, r, "/authentification", http.StatusSeeOther)
}

// -------------------------------------------- Ajouter données d'un post dans la DB

func doPost(w http.ResponseWriter, r *http.Request) {
	userInfo := GetDataUserToken(r)
	var Post POST
	var UserPost []POST
	if !CheckCookie(r) {
		http.Redirect(w, r, "/authentification", http.StatusSeeOther)
		return
	}
	if r.Method == "POST" {
		Post.Post_Name = r.FormValue("post-title")
		Post.Content = r.FormValue("post-content")
		AddPost(Post.Post_Name, Post.Content, r, 1)
		URL := "/post?id="
		UserPost = GetPostUser(strconv.Itoa(userInfo.Id))
		lastPostUser := UserPost[len(UserPost)-1]
		lastPostUserID := lastPostUser.ID
		URL += strconv.Itoa(lastPostUserID)
		http.Redirect(w, r, URL, http.StatusSeeOther)
	} else {
		data := HTML{Users: userInfo, PostsUser: UserPost, Post: Post}
		files := []string{"../static/post.html", "../static/navbar.html"}
		tmpl := template.Must(template.ParseFiles(files...))
		tmpl.Execute(w, data)
	}
}

// -------------------------------------------- AUTRE

func UploadFile(w http.ResponseWriter, r *http.Request) { // Ajout d'une photo de profile pour le User
	userInfo := GetDataUserToken(r)
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("uploadFile")
	checkErr(err)
	defer file.Close()
	fmt.Printf("File Size: %+v\n", handler.Size)
	tempFile, err := os.Create("../assets/users/upload-" + userInfo.Username + ".png")
	checkErr(err)
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	checkErr(err)
	tempFile.Write(fileBytes)
	AddUserPic("username", userInfo.Username)
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}
