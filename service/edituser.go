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


func EditUsername(ctx *gin.Context) {
	newname, _ := ctx.GetPostForm("name")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Get now user name
	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
    if err != nil {
        ctx.String(http.StatusOK, "cookie is nil")
		return
    }else{
		fmt.Println(name, "を変更します")
	}

	//check double name
	var user database.User
	err = db.QueryRow("SELECT * FROM users WHERE name = ?", newname).Scan(&user.Name, &user.Password)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("この名前でアカウントを作成できます")
	case err != nil:
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	default:
		ctx.HTML(http.StatusOK, "user_edit.html", gin.H{"Title": "USER IOFO","Info":"このアカウント名は既に使われています"})
		return
	}

	// Change task in DB
	data := map[string]interface{}{"newname" : newname, "name": name}
	res , err := db.NamedExec("UPDATE users SET name = (:newname) WHERE name=(:name)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Change task in DB
	res , err = db.NamedExec("UPDATE tasks SET name = (:newname) WHERE name=(:name)", data)
		if res == nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
	
		// Set Cookie
	ctx.SetCookie("name", newname, 1000, "/", "localhost", false, true)
	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME","Head2": "変更完了", "Name":name, "Done":"さんがユーザー名を"+newname+"に変更しました"})
}

func EditPassword(ctx *gin.Context) {
	newpass, _ := ctx.GetPostForm("password")
	hashed_pass := sha256.Sum256([]byte(newpass))
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Get now user name
	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
    if err != nil {
        ctx.String(http.StatusOK, "cookie is nil")
		return
    }else{
		fmt.Println(name, "のパスワードを変更します")
	}

	// Change task in DB
	data := map[string]interface{}{"newpass" : hex.EncodeToString(hashed_pass[:]), "name": name}
	res , err := db.NamedExec("UPDATE users SET password = (:newpass) WHERE name=(:name)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME" ,"Head2": "変更完了", "Name":name, "Done":"さんがパスワードを変更しました"})
}