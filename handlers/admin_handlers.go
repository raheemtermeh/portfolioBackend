package handlers

import (
	"net/http"
	"portfolioBackend/database"
	"portfolioBackend/models"

	"github.com/gin-gonic/gin"
)

// --- CRUD Operations for Projects ---

// GetProjects تمام پروژه‌ها را برمی‌گرداند
func GetProjects(c *gin.Context) {
	var projects []models.Project
	database.DB.Order("display_order asc").Find(&projects)
	c.JSON(http.StatusOK, projects)
}

// CreateProject یک پروژه جدید می‌سازد
func CreateProject(c *gin.Context) {
	var input models.Project
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}

// UpdateProject یک پروژه موجود را به‌روزرسانی می‌کند
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found!"})
		return
	}
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&project)
	c.JSON(http.StatusOK, project)
}

// DeleteProject یک پروژه را حذف می‌کند
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	if err := database.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found!"})
		return
	}
	database.DB.Delete(&project)
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully!"})
}