package models

import (
   
    "time"
)

type Fact struct {
    ID          uint           `gorm:"primaryKey"`
    Title       string         `gorm:"text;not null;default:null"`
    Description string         `gorm:"text;not null;default:null"`
    ImageURI    string         `gorm:"text;not null;default:null"`
    CreatedDate time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

func (Fact) TableName() string {
    return "tb_casestudy"
}
