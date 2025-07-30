package models

import "time"

type TimelineEvent struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    IconName   string    `json:"icon_name"`
    Title      string    `json:"title"`
    Subtitle   string    `json:"subtitle"`
    IsLeft     bool      `json:"is_left"`
    EventDate  time.Time `json:"event_date"`
}