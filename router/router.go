package router

import (
	"example/employee-manager/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Router(dbConn *gorm.DB) {
	managerHandler := controllers.ManagerHandler{
		Db: dbConn,
	}
	r := gin.Default()
	r.LoadHTMLGlob("views/html/*.html")
	r.GET("/manager", managerHandler.GetAll)
	r.Run(":9000")
}
