package models

type Level struct {
	LevelID   int64      `json:"levelID" gorm:"primaryKey"`
	LevelName int64      `json:"levelName" gorm:"<-;unique;not null"`
	Questions []Question `json:"questions" gorm:"foreignKey:LevelID"`
	Users     []User     `json:"users" gorm:"foreignKey:LevelID"`
}
