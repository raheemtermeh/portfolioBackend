package handlers

import (
	"net/http"
	"portfolioBackend/database"
	"portfolioBackend/models"

	"github.com/gin-gonic/gin"
)

// GetAllContent تمام محتوای مورد نیاز برای صفحه اصلی را برمی‌گرداند
func GetAllContent(c *gin.Context) {
	var projects []models.Project
	var skills []models.Skill
	var timelineEvents []models.TimelineEvent
	var siteConfigs []models.SiteConfig

	// دریافت تمام اطلاعات با یک بار کوئری به دیتابیس
	database.DB.Order("display_order asc").Find(&projects)
	database.DB.Order("display_order asc").Find(&skills)
	database.DB.Order("event_date asc").Find(&timelineEvents)
	database.DB.Find(&siteConfigs)

	// تبدیل تنظیمات سایت به یک map برای دسترسی راحت‌تر در فرانت‌اند
	configMap := make(map[string]string)
	for _, config := range siteConfigs {
		configMap[config.Key] = config.Value
	}

	c.JSON(http.StatusOK, gin.H{
		"configs":  configMap,
		"projects": projects,
		"skills":   skills,
		"timeline": timelineEvents,
	})
}