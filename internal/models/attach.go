package models

type AttachBanner struct {
	BannerID int32 `json:"banner_id,omitempty"`
	SlotID   int32 `json:"slot_id,omitempty"`
}
