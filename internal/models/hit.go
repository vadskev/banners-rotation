package models

type HitBanner struct {
	BannerID      int32 `json:"banner_id,omitempty"`
	SlotID        int32 `json:"slot_id,omitempty"`
	SocialGroupID int32 `json:"social_group_id,omitempty"`
}
