package models

type SocialMediaPlatform struct {
	BaseModel
	Name    string `json:"name" gorm:"unique;not null"`
	LogoURL string `json:"logo"`
	Socials []UserSocial
}

func init() {
	Migration = append(Migration, &SocialMediaPlatform{})
}
