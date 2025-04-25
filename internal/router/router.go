package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"learn_verse/internal/handler"
	"learn_verse/internal/repository"
	"learn_verse/internal/service"
)

func Setup(r *gin.Engine, db *sql.DB) {
	// Repos
	collRepo := repository.NewCollectionRepo(db)
	resRepo := repository.NewResourceRepo(db)

	// Services
	collSvc := service.NewCollectionService(collRepo)
	resSvc := service.NewResourceService(resRepo)

	// Handlers
	collH := handler.NewCollectionHandler(collSvc)
	resH := handler.NewResourceHandler(resSvc)

	// Routes
	api := r.Group("/api")
	api.POST("/collections", collH.Create)
	api.GET("/collections", collH.List)
	api.GET("/collections/:id", collH.Get)
	api.DELETE("/collections/:id", collH.Delete)
	api.PUT("/collections/:id", collH.Update)
	api.POST("/resources", resH.Create)

}
