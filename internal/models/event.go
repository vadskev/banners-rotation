package models

type Message struct {
	Type          string `json:"type"`
	SlotID        int32  `json:"slot_id"`
	BannerID      int32  `json:"banner_id"`
	SocialGroupID int32  `json:"social_group_id"`
	CreatedAt     int64  `json:"createdAt"`
}
