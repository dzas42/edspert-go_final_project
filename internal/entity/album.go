package entity

// Album is  entity will be used for album definition
type Album struct {
	ID       int64   `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	Title    string  `json:"title"`
	ArtistId *int64  `json:"artist_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Artist   *Artist `json:"artist,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
	Price    float32 `json:"price"`
}
