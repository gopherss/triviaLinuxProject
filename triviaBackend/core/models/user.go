package models

type User struct {
	UserID       int64  `json:"userID" gorm:"primaryKey"`
	Name         string `json:"name" gorm:"<-"`
	LastName     string `json:"lastName" gorm:"<-"`
	NickName     string `json:"nickName" gorm:"<-"`
	Type         string `json:"type" gorm:"<-"`
	Email        string `json:"email" gorm:"<-"`
	Password     string `json:"password" gorm:"<-"`
	CurrentScore int64  `json:"currentScore" gorm:"<-"`
	LevelID      int64  `json:"levelID"`
}
