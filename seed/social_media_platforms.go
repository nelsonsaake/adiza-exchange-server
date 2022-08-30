package seed

import (
	"projects/adiza-exchange-server/archetypes"
	"projects/adiza-exchange-server/repo"
)

func UpdateRepoSocialMediaPlatforms() {
	for _, platform := range archetypes.SocialMediaPlatforms {
		repo.DB.Save(platform)
	}
}

func init() {
	UpdateRepoSocialMediaPlatforms()
}
