package models

type Scan struct {
	BaseModel
	UserID        uint `json:"userId"`
	User          User `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ScannedUserID uint `json:"scannedUserID"`
	ScannedUser   User `json:"scannedUser" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func init() {
	Migration = append(Migration, &Scan{})
}
