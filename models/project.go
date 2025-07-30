package models

type Project struct {
    ID           uint   `json:"id" gorm:"primary_key"`
    Name         string `json:"name"`
    Description  string `json:"description"`
    URL          string `json:"url"`
    Icon         string `json:"icon"`
    DisplayOrder int    `json:"display_order"`
}