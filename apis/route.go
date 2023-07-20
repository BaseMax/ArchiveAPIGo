package apis

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo) {

	UG := e.Group("/u")
	UG.POST("/login", Login)
	UG.POST("/register", Register)
	UG.GET("/users/:id", GetUser)
	UG.PUT("/users/:id", UpdateUser)
	UG.DELETE("/users/:id", DeleteUser)

	CG := e.Group("/c")
	CG.GET("/categories", GetCategories)
	CG.GET("/categories/:id", GetCategoryByID)
	CG.POST("/categories", CreateCategory)
	CG.PUT("/categories/:id", UpdateCategory)
	CG.DELETE("/categories/:id", DeleteCategory)

	FG := e.Group("/f")
	FG.GET("/files", GetFiles)
	FG.GET("/files/:id", GetFileById)
	FG.POST("/files", UploadFile)
	FG.PUT("/files/:id", UpdateFile)
	FG.DELETE("/files/:id", DeleteFile)
	FG.GET("/files/public-link/:id", GetPublicLink)
	FG.GET("/search/:keyword", SearchFile)
}
