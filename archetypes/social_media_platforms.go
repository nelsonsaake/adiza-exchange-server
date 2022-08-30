package archetypes

import "projects/adiza-exchange-server/models"

var (
	FaceBook = models.SocialMediaPlatform{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Name:    "facebook",
		LogoURL: "/logos/facebook.svg",
	}

	Twitter = models.SocialMediaPlatform{
		BaseModel: models.BaseModel{
			ID: 2,
		},
		Name:    "twitter",
		LogoURL: "/logos/twitter.svg",
	}

	LinkedIn = models.SocialMediaPlatform{
		BaseModel: models.BaseModel{
			ID: 3,
		},
		Name:    "linkedIn",
		LogoURL: "/logos/linedIn.svg",
	}

	Instagram = models.SocialMediaPlatform{
		BaseModel: models.BaseModel{
			ID: 4,
		},
		Name:    "instagram",
		LogoURL: "/logos/instagram.svg",
	}

	WhatsApp = models.SocialMediaPlatform{
		BaseModel: models.BaseModel{
			ID: 5,
		},
		Name:    "whatsapp",
		LogoURL: "/logos/whatsapp.svg",
	}

	SocialMediaPlatforms = map[uint]*models.SocialMediaPlatform{
		FaceBook.ID:  &FaceBook,
		Twitter.ID:   &Twitter,
		LinkedIn.ID:  &LinkedIn,
		Instagram.ID: &Instagram,
		WhatsApp.ID:  &WhatsApp,
	}
)
