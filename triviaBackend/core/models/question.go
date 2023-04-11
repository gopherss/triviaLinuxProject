package models

type Question struct {
	QuestionID  int64    `json:"questionID" gorm:"primaryKey"`
	Ask         string   `json:"ask" gorm:"<-"`
	Reply       string   `json:"reply" gorm:"<-"`
	Score       int64    `json:"score" gorm:"<-"`
	Information string   `json:"information" gorm:"<-"`
	LevelID     int64    `json:"levelID"`
	Options     []Option `json:"options" gorm:"foreignKey:QuestionID"`
}
