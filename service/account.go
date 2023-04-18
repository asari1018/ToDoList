package service

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/url"
	"fmt"

	"github.com/gin-gonic/gin"
	database "todolist.go/db"
	"database/sql"
)

func NewAccount(ctx *gin.Context) {
	name, _ := ctx.GetPostForm("name")
	password, _ := ctx.GetPostForm("password")
	hashed_pass := sha256.Sum256([]byte(password))

	data := map[string]interface{}{"name":name, "password":hex.EncodeToString(hashed_pass[:])}

	if name=="" || hex.EncodeToString(hashed_pass[:])=="" {
		ctx.HTML(http.StatusOK, "first.html", gin.H{"Title": "REGISTERED","Info":"未記入のフォームがあります"})
		return
	}

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	//check double name
	var user database.User
	err = db.QueryRow("SELECT * FROM users WHERE name = ?", name).Scan(&user.Name, &user.Password)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("この名前でアカウントを作成できます")
	case err != nil:
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	default:
		ctx.HTML(http.StatusOK, "first.html", gin.H{"Title": "REGISTERED","Info":"このアカウント名は既に使われています"})
		return
	}
	
	// Set Cookie
	ctx.SetCookie("name", name, 1000, "/", "localhost", false, true)
	res , err := db.NamedExec("INSERT INTO users (name, password) VALUES (:name, :password)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME", "Name":name, "Done":"さんのアカウントが作成されましました"})
}

func Certify(ctx *gin.Context) {
	name, _ := ctx.GetPostForm("name")
	password, _ := ctx.GetPostForm("password")
	hashed_pass := sha256.Sum256([]byte(password))

	if name=="" || password=="" {
		ctx.HTML(http.StatusOK, "login.html", gin.H{"Title": "LOGIN","Info":"未記入のフォームがあります"})
		return
	}

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Serch Accoumt
	var user database.User
	err = db.QueryRow("SELECT * FROM users WHERE name = ?", name).Scan(&user.Name, &user.Password)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("存在しません")
		ctx.HTML(http.StatusOK, "login.html", gin.H{"Title": "LOGIN", "Info":"ユーザー名またはパスワードが間違っています"})
		return
	case err != nil:
		ctx.String(http.StatusInternalServerError, err.Error())
	default:
		fmt.Println(user.ID, user.Name, user.Password)
	}
	// Certify Account
	if hex.EncodeToString(hashed_pass[:]) == user.Password{
		// Set Cookie
		ctx.SetCookie("name", name, 1000, "/", "localhost", false, true)
		ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME", "Head2":"ようこそ", "Name":name, "Done":"さんがログインしました"})
		return
	} else{
		ctx.HTML(http.StatusOK, "login.html", gin.H{"Title": "LOGIN", "Info":"ユーザー名またはパスワードが間違っています"})
		return
	}
}

func DeleteUser(ctx *gin.Context) {
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
    if err != nil {
        ctx.String(http.StatusOK, "cookie is nil")
		return
    }

	// Delete user in DB
	data := map[string]interface{}{"name": name}
	res , err := db.NamedExec("DELETE FROM users WHERE name=(:name)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	res , err = db.NamedExec("DELETE FROM tasks WHERE name=(:name)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "login.html", gin.H{"Title": "LOGIN" ,"Info":name+"を削除しました"})
}

func Logout(ctx *gin.Context) {
	// Set Cookie

	cookie, err := ctx.Request.Cookie("name")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
    ctx.SetCookie("name", cookie.Value, -1, "/", "localhost", false, true)
	ctx.HTML(http.StatusOK, "login.html", gin.H{"Title": "LOGIN" ,"Info":"ログアウトしました"})
}