package entities

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/guregu/null.v4"
)

type Student struct {
	Model
	Login        string      `json:"login" valid:"required,type(string),length(5|255)" gorm:"size:255"`
	Password     string      `json:"password" valid:"required,type(string),length(8|255)" gorm:"size:255"`
	Token        string      `json:"token" gorm:"size:255"`
	FirstName    string      `json:"first_name" valid:"required,type(string),length(2|255)" gorm:"size:255"`
	LastName     string      `json:"last_name" valid:"required,type(string),length(2|255)" gorm:"size:255"`
	Avatar       string      `json:"avatar" valid:"type(string),length(1|255)" gorm:"size:255"`
	Status       string      `json:"status" valid:"type(string),length(1|255)" gorm:"size:255"`
	Patronymic   string      `json:"patronymic" valid:"type(string),length(2|255),optional" gorm:"size:255"`
	StudyGroupId null.Int    `json:"study_group_id"`
	StudyGroup   *StudyGroup `json:"study_group"`
}

func (student *Student) TableName() string {
	return "students"
}

type UserClaims struct {
	Authorized bool
	Student
	jwt.StandardClaims
}

func (student *Student) GenerateJWT() error {
	var signingKey = []byte("aumsu-portal-backend")
	claims := UserClaims{
		Authorized: true,
		Student:    *student,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	student.Token = tokenString

	return err
}
