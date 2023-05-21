package app

import (
	"hook/bootstrap"
)

func Init() {
	OPTIONS = bootstrap.ReadCliOption()
}

func Boot() {
	CFG = bootstrap.LoadConfig(OPTIONS.Config)

	bootstrap.InitLogger(OPTIONS.Config)
	GIN = bootstrap.InitGin()
}
