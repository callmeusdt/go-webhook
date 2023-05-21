package app

import (
	"github.com/gin-gonic/gin"
	"hook/bootstrap"
)

var GIN *gin.Engine
var CFG *bootstrap.Config
var OPTIONS *bootstrap.CliOption
