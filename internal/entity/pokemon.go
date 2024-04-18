package entity

type Pokemon struct {
	ID   int64  `json:"id" db:"id"`
	Name string `xml:"name" db:"name" binding:"required"`
}
