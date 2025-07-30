package handlers

import (
	"net/http"
	"portfolioBackend/database"
	"portfolioBackend/models"

	"github.com/gin-gonic/gin"
)

// CreateContactMessage یک پیام جدید از فرم تماس دریافت و ذخیره می‌کند
func CreateContactMessage(c *gin.Context) {
	var input models.ContactMessage

	// اتصال (Bind) JSON ورودی به مدل
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ایجاد رکورد جدید در دیتابیس
	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Your message has been sent successfully!"})
}
// GetContactMessages تمام پیام‌های تماس را از دیتابیس دریافت می‌کند
func GetContactMessages(c *gin.Context) {
    var messages []models.ContactMessage

    // دریافت پارامتر جستجو از query string
    searchQuery := c.Query("search")

    // ساخت کوئری پایه
    query := database.DB.Order("created_at desc")

    // اگر پارامتر جستجو وجود داشت، شرط WHERE اضافه می‌شود
    if searchQuery != "" {
        searchTerm := "%" + searchQuery + "%"
        query = query.Where("name LIKE ? OR email LIKE ?", searchTerm, searchTerm)
    }

    if err := query.Find(&messages).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
        return
    }

    c.JSON(http.StatusOK, messages)
}

func DeleteContactMessage(c *gin.Context) {
    // گرفتن ID از پارامترهای مسیر
    id := c.Param("id")

    // پیدا کردن پیام با ID مورد نظر
    var message models.ContactMessage
    if err := database.DB.First(&message, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Message not found!"})
        return
    }

    // حذف پیام از دیتابیس
    database.DB.Delete(&message)

    c.JSON(http.StatusOK, gin.H{"message": "Message deleted successfully!"})
}

func MarkAsRead(c *gin.Context) {
    id := c.Param("id")
    var message models.ContactMessage
    if err := database.DB.First(&message, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Message not found!"})
        return
    }

    message.Read = true
    database.DB.Save(&message)

    c.JSON(http.StatusOK, message)
}

func GetDashboardStats(c *gin.Context) {
    var totalMessages int64
    var unreadMessages int64

    database.DB.Model(&models.ContactMessage{}).Count(&totalMessages)
    database.DB.Model(&models.ContactMessage{}).Where("read = ?", false).Count(&unreadMessages)

    c.JSON(http.StatusOK, gin.H{
        "total_messages": totalMessages,
        "unread_messages": unreadMessages,
    })
}