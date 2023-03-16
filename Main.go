package main

import (
	"Tugas-Pert4/config"
	"Tugas-Pert4/controller"
	"Tugas-Pert4/midleware"
	"Tugas-Pert4/repository"
	"Tugas-Pert4/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database := config.SetDatabaseConnection()
	defer config.CloseDatabaseConnection(database)

	server := gin.Default()
	// middleware CORS
	server.Use(midleware.CORSMiddleware())
	// server.Use(midleware.Authenticate())

	userRepo := repository.NewUserRepository(database)
	blogRepo := repository.NewBlogRepository(database)
	commentRepo := repository.NewCommentRepository(database)

	req_user := service.NewUserHandler(userRepo)
	service.NewUserHandler(userRepo)
	req_blog := service.NewBlogHandler(blogRepo)
	req_comment := service.NewCommentHandler(commentRepo)

	ctrl_user := controller.NewUserController(req_user)
	ctrl_blog := controller.NewBlogController(req_blog)
	ctrl_comment := controller.NewCommentController(req_comment)

	server.POST("/signup", ctrl_user.AddAccount)
	server.POST("/login", ctrl_user.LoginAccount)
	server.GET("/validate/:id", ctrl_user.ValidateAccount)
	server.GET("/user", ctrl_user.GetAccount)
	server.PATCH("/userprofil/:id", ctrl_user.UpdateAccount)
	server.GET("/userprofil/:id", ctrl_user.GetAccountByID)
	server.DELETE("/user/:id", ctrl_user.DeleteAccount)

	server.POST("/user/blog", ctrl_blog.AddBlog)
	server.GET("/user/blog", ctrl_blog.GetBlog)
	server.GET("/user/blog/detail", ctrl_blog.GetBlogDetail)
	server.DELETE("/user/blog/:id", ctrl_blog.DeleteBlog)
	server.DELETE("/user/blog/user/:id", ctrl_blog.DeleteBlogUser)
	server.PATCH("/user/blog/:id", ctrl_blog.UpdateBlog)

	server.POST("/user/blog/comment", ctrl_comment.AddComment)
	server.GET("/user/blog/comment", ctrl_comment.GetComment)
	server.DELETE("/user/blog/comment/:id", ctrl_comment.DeleteComment)
	server.PATCH("/user/blog/comment/:id", ctrl_comment.UpdateComment)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run("localhost:" + port)
}
