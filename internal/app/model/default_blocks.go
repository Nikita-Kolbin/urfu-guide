package model

type DefaultBlock struct {
	ID          int    `db:"id" json:"-"`
	SectionID   int    `db:"section_id" json:"-"`
	Address     string `db:"address" json:"address"`
	Timetable   string `db:"timetable" json:"timetable"`
	Phone       string `db:"phone" json:"phone"`
	Description string `db:"description" json:"description"`
	Link        string `db:"link" json:"link"`
	Image       string `db:"image" json:"image"`
	Position    int    `db:"position" json:"-"`
}

type DefaultBlockList []*DefaultBlock
