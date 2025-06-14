package model

type DropDownBlock struct {
	ID        int                 `db:"id" json:"-"`
	SectionID int                 `db:"section_id" json:"-"`
	Title     string              `db:"title" json:"title"`
	SubTitle  string              `db:"sub_title" json:"sub_title"`
	Icon      string              `db:"icon" json:"icon"`
	Position  int                 `db:"position" json:"-"`
	Elements  DropDownElementList `json:"elements"`
}

type DropDownBlockList []*DropDownBlock

type DropDownElement struct {
	ID          int    `db:"id" json:"-"`
	BlockID     int    `db:"block_id" json:"-"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Link        string `db:"link" json:"link"`
	Position    int    `db:"position" json:"-"`
}

type DropDownElementList []*DropDownElement
