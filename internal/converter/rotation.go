package converter

import (
	"github.com/vadskev/banners-rotation/internal/models"
	desc "github.com/vadskev/banners-rotation/pkg/rotation_v1"
)

func ToBannerFromAddBannerRequest(info *desc.AddBannerRequest) *models.Banner {
	return &models.Banner{
		Description: info.Description,
	}
}
func ToBannerFromDeleteBannerRequest(info *desc.DeleteBannerRequest) *models.Banner {
	return &models.Banner{
		ID: info.Id,
	}
}

func ToSlotFromAddSlotRequest(info *desc.AddSlotRequest) *models.Slot {
	return &models.Slot{
		Description: info.Description,
	}
}
func ToSlotFromDeleteSlotRequest(info *desc.DeleteSlotRequest) *models.Slot {
	return &models.Slot{
		ID: info.Id,
	}
}

func ToSocialGroupFromAddSocialGroupRequest(info *desc.AddSocialGroupRequest) *models.SocialGroup {
	return &models.SocialGroup{
		Description: info.Description,
	}
}
func ToSocialGroupFromSocialGroupRequest(info *desc.DeleteSocialGroupRequest) *models.SocialGroup {
	return &models.SocialGroup{
		ID: info.Id,
	}
}
