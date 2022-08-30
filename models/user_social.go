package models

type UserSocial struct {
	BaseModel
	UserID                uint   `json:"userId"`
	SocialMediaPlatformID uint   `json:"socialMediaPlatformId" gorm:"UniqueIndex:UniqueSocialMediaHandle;not null"`
	Handle                string `json:"handle" gorm:"UniqueIndex:UniqueSocialMediaHandle;not null"`
}

func init() {
	Migration = append(Migration, &UserSocial{})
}
