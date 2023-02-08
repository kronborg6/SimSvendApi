package main

import (
	"github.com/kronborg6/SimSvendApi/api/db"
	"github.com/kronborg6/SimSvendApi/api/models"
)

func main() {
	db := db.Init()

	models.Setup(db)
}
