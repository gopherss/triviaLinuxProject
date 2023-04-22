package models

type Option struct {
	OptionID   int64  `json:"optionID" gorm:"primaryKey"`
	One        string `json:"one" gorm:"<-"`
	Two        string `json:"two" gorm:"<-"`
	Three      string `json:"three" gorm:"<-"`
	QuestionID int64  `json:"questionID"`
}
