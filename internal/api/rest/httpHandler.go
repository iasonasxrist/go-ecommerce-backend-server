package rest

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RestHandler struct {
	App *gin.Engine
	Db *gorm.DB
}
