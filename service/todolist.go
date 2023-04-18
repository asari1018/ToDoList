package service

import (
	"net/http"
	"strconv"
	"net/url"
	"fmt"

	"github.com/gin-gonic/gin"
	database "todolist.go/db"
)


// TaskList renders list of tasks in DB
func TaskList(ctx *gin.Context) {

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		fmt.Println("データへのアクセスに失敗")
		ctx.String(http.StatusInternalServerError, err.Error())
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

	// Get tasks in DB
	var tasks []database.Task
	err = db.Select(&tasks, "SELECT * FROM tasks WHERE name = ?", name) // Use DB#Select for multiple entries
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Render tasks
	ctx.HTML(http.StatusOK, "task_list.html", gin.H{"Title": "Task list", "Tasks": tasks})
}

// TrueList renders true list of tasks in DB
func TrueList(ctx *gin.Context) {
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
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

	// Get true tasks in DB
	var tasks []database.Task
	err = db.Select(&tasks, "SELECT * FROM tasks WHERE name = ? AND is_done = true ", name) // Use DB#Select for multiple entries
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Render tasks
	ctx.HTML(http.StatusOK, "task_list.html", gin.H{"Title": "Task list", "Tasks": tasks})
}

// False List renders list of tasks in DB
func FalseList(ctx *gin.Context) {
	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	//クッキーの名前からタスクのIDを得る
	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
    if err != nil {
        ctx.String(http.StatusOK, "cookie is nil")
		return
    }

	// Get tasks in DB
	var tasks []database.Task
	err = db.Select(&tasks, "SELECT * FROM tasks WHERE name = ? AND is_done = false", name) // Use DB#Select for multiple entries
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Render tasks
	ctx.HTML(http.StatusOK, "task_list.html", gin.H{"Title": "Task list", "Tasks": tasks})
}


// TaskAdd renders new tasks in DB
func TaskAdd(ctx *gin.Context) {
	title, _ := ctx.GetPostForm("title")

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
    }else{
		fmt.Println(name, "がタスクを登録します")
	}

	// Add tasks in DB
	data := map[string]interface{}{"name":name, "title":title}
	res , err := db.NamedExec("INSERT INTO tasks (name, title) VALUES (:name, :title)", data)
	if res == nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Render tasks
	ctx.HTML(http.StatusOK, "welcome.html", gin.H{"Title": "HOME","Head2": "タスク登録完了", "Name":title, "Done":"を登録しました"})
}

// ShowTask renders a task with given ID
func ShowTask(ctx *gin.Context) {

	// Get DB connection
	db, err := database.GetConnection()
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// parse ID given as a parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	cookie, err := ctx.Request.Cookie("name")
	name, _ := url.QueryUnescape(cookie.Value)
    if err != nil {
        ctx.String(http.StatusOK, "cookie is nil")
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

