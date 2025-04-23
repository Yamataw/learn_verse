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
	grp := r.Group("/api")
	grp.POST("/collections", collH.Create)
	grp.POST("/resources", resH.Create)
}
