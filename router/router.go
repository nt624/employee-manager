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
	router := gin.Default()
	router.LoadHTMLGlob("views/html/*.html")
	router.GET("/manager", managerHandler.GetAll)

}
