package models

import "golang.org/x/crypto/bcrypt"

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

func (u *User) HashPassword() (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	return string(bytes), err
}

func (u *User) CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil

}
