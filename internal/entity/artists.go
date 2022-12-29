package entity

type Artist struct {
	ID     int64   `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	Name   string  `json:"name"`
	Albums []Album `json:"artist,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
