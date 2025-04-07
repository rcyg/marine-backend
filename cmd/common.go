package cmd

import (
	"marine-backend/internal/bootstrap"
	"marine-backend/internal/db"
)

func Init() {
	bootstrap.InitConfig()
	bootstrap.Log()
	bootstrap.InitDB()
}

func Release() {
	db.Close()
}
