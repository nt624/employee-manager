package controllers

import (
	"example/employee-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ManagerHandler struct {
	Db *gorm.DB
}

func (h *ManagerHandler) GetAll(c *gin.Context) {
	var emp []models.Employee
	err := h.Db.Select("id", "name", "dept", "age", "email", "tel", "phone").Find(&emp)
	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"emp": emp,
	})
}
