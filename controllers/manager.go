package controllers

import (
	"example/employee-manager/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ManagerHandler struct {
	Db *gorm.DB
}

func (h *ManagerHandler) GetAll(c *gin.Context) {
	var emp []models.Employee
	err := h.Db.Find(&emp).Error
	if err != nil {
		panic(err.Error())
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"emp": emp,
	})
}

func (h *ManagerHandler) GetOne(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		panic(err.Error())
	}
	var emp []models.Employee
	err = h.Db.First(&emp, id).Error
	if err != nil {
		panic(err.Error())
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"emp": emp,
	})
}
