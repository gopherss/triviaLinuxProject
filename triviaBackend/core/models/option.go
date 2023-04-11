package models

type Option struct {
	OptionID   int64  `json:"optionID" gorm:"primaryKey"`
	One        string `json:"one" gorm:"<-"`
	Two        string `json:"Two" gorm:"<-"`
	Tree       string `json:"Tree" gorm:"<-"`
	QuestionID int64  `json:"questionID"`
}
