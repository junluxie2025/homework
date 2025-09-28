package routers

import (
	"homework/go-basic/task4/controllers"
	"homework/go-basic/task4/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.New()

	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.ErrorHandlerMiddleware())
	r.Use(gin.Recovery())

	authController := &controllers.AuthController{}
	postController := &controllers.PostController{}
	commentController := &controllers.CommentController{}

	api := r.Group("api/v1")
	{
		//认证
		auth := api.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
		}

		//中间件认证
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			//用户信息
			authenticated.GET("/profile", authController.GetProfile)

			//文章
			posts := authenticated.Group("/posts")
			{
				posts.POST("", postController.CreatePost)
				posts.PUT("/:id", postController.UpdatePost)
				posts.DELETE("/:id", postController.DeletePost)
			}

			//评论
			comments := authenticated.Group("/posts/:post_id/comments")
			{
				comments.POST("", commentController.CreateComment)
			}
		}

		//公开
		public := api.Group("")
		{
			public.GET("/posts", postController.GetPosts)
			public.GET("/posts/:id", postController.GetPost)
		}

		//评论
		comments := api.Group("/comments")
		{
			comments.GET("/post/:post_id", commentController.GetComments)
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Blog API is running",
		})
	})

	return r
}
