package models

type AttachBanner struct {
	ID       int32 `json:"id,omitempty"`
	BannerID int32 `json:"banner_id,omitempty"`
	SlotID   int32 `json:"slot_id,omitempty"`
}
