package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{"Title": "LOGIN"})
}

func First(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "first.html", gin.H{"Title": "REGISTER"})
}

func Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "task_add.html", gin.H{"Title": "ADD"})
}

func EditUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user_edit.html", gin.H{"Title": "USER INFO"})
}


