package rest

import (
	"ecommerce.com/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RestHandler struct {
	App  *gin.Engine
	Db   *gorm.DB
	Auth helper.Auth
}
