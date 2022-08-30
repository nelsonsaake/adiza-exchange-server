package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Photo        string       `json:"photo" `
	FirstName    string       `json:"firstName"`
	LastName     string       `json:"lastName"`
	OtherNames   string       `json:"otherNames"`
	JobTitle     string       `json:"jobTitle"`
	Email        string       `json:"email" gorm:"unique;not null"`
	PhoneNumber  string       `json:"phoneNumber" gorm:"unique;not null"`
	Password     string       `json:"-"`
	Socials      []UserSocial `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ScannedUsers []Scan       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ScannedBy    []Scan       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	tx.Where("user_id = ?", u.ID).Delete(&UserSocial{})
	tx.Where("user_id = ?", u.ID).Delete(&Scan{})
	return
}

func (u *User) String() string {
	buf := bytes.Buffer{}
	b, _ := json.MarshalIndent(u, "", "\t")
	buf.WriteString(fmt.Sprintf("%s\n", b))
	buf.WriteString(fmt.Sprintln("ID: ", u.BaseModel.ID))
	buf.WriteString(fmt.Sprintln("Created At: ", u.BaseModel.CreatedAt))
	buf.WriteString(fmt.Sprintln("Updated At: ", u.BaseModel.UpdatedAt))
	buf.WriteString(fmt.Sprintln("Deleted At: ", u.BaseModel.DeletedAt))
	return buf.String()
}

func init() {
	Migration = append(Migration, &User{})
}
