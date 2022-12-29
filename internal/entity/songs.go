package entity

type Song struct {
	ID      int64  `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	AlbumId *int64 `json:"album_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Album   *Album `json:"album,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title   string `json:"title"`
	Lyrics  string `json:"lyrics"`
}
