package models

type SiteConfig struct {
    Key   string `json:"key" gorm:"primary_key"`
    Value string `json:"value"`
}