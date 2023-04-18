package service

import (
	"net/http"
	"strconv"
	"net/url"
	"fmt"

	"github.com/gin-gonic/gin"
	database "todolist.go/db"
)

func ChangeTitle(ctx *gin.Context) {
	strid, _ := ctx.GetPostForm("id")
	title, _ := ctx.GetPostForm("title")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// parse ID given as a parameter
	id, err := strconv.Atoi(strid)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	//クッキーの名前からタスクのIDを得る
	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
    if err != nil {
        ctx.String(http.StatusOK, "cookie is nil")
		return
    }else{
		fmt.Println(name, "のタスクを表示します")
	}


	// Change task in DB
	data := map[string]interface{}{"title" : title, "id": id, "name":name}
	res , err := db.NamedExec("UPDATE tasks SET title = (:title) WHERE id=(:id) AND name=(:name)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Get a task with given ID
	var task database.Task
	err = db.Get(&task, "SELECT * FROM tasks WHERE id=?", id) // Use DB#Get for one entry
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if task.Name == name{
		ctx.HTML(http.StatusOK, "task_edit.html", gin.H{"task": task})
		return
	} else {
		ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME","Head2": "アクセス権がありません", "Name":id, "Done":"　このタスクへのアクセス権がありません"})
		return
	}
}


func TaskToTrue(ctx *gin.Context) {
	strid, _ := ctx.GetPostForm("id")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// parse ID given as a parameter
	id, err := strconv.Atoi(strid)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	//クッキーの名前からタスクのIDを得る
	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
    if err != nil {
        ctx.String(http.StatusOK, "cookie is nil")
		return
    }else{
		fmt.Println(name, "のタスクを表示します")
	}

	// Change task in DB
	data := map[string]interface{}{"id": id, "name":name}
	res , err := db.NamedExec("UPDATE tasks SET is_done = true WHERE id=(:id) AND name=(:name)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Get a task with given ID
	var task database.Task
	err = db.Get(&task, "SELECT * FROM tasks WHERE id=?", id) // Use DB#Get for one entry
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if task.Name == name{
		ctx.HTML(http.StatusOK, "task_edit.html", gin.H{"task": task})
		return
	} else {
		ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME","Head2": "アクセス権がありません", "Name":id, "Done":"　このタスクへのアクセス権がありません"})
		return
	}
}


func TaskToFalse(ctx *gin.Context) {
	strid, _ := ctx.GetPostForm("id")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// parse ID given as a parameter
	id, err := strconv.Atoi(strid)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	//クッキーの名前からタスクのIDを得る
	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
	if err != nil {
		ctx.String(http.StatusOK, "cookie is nil")
		return
	}else{
		fmt.Println(name, "のタスクを表示します")
	}


	// Change task in DB
	data := map[string]interface{}{"id": id, "name":name}
	res , err := db.NamedExec("UPDATE tasks SET is_done = false WHERE id=(:id) AND name=(:name)", data)
		if res == nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
	
	// Get a task with given ID
	var task database.Task
	err = db.Get(&task, "SELECT * FROM tasks WHERE id=?", id) // Use DB#Get for one entry
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	if task.Name == name{
		ctx.HTML(http.StatusOK, "task_edit.html", gin.H{"task": task})
		return
	} else {
		ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME","Head2": "アクセス権がありません", "Name":id, "Done":"　このタスクへのアクセス権がありません"})
		return
	}

}



func TaskDelete(ctx *gin.Context) {
	strid, _ := ctx.GetPostForm("id")

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// parse ID given as a parameter
	id, err := strconv.Atoi(strid)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	//クッキーの名前からタスクのIDを得る
	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
	if err != nil {
		ctx.String(http.StatusOK, "cookie is nil")
		return
	}else{
		fmt.Println(name, "のタスクを表示します")
	}

	// Delete task in DB
	data := map[string]interface{}{"id": id, "name":name}
	res , err := db.NamedExec("DELETE FROM tasks WHERE id=(:id) AND name=(:name)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME", "Done":"タスクの消去が完了しました"})
}
