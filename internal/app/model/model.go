package model

type Language struct {
	Code     string      `db:"code" json:"code"`
	Sections SectionList `json:"sections"`
}

type LanguageList []*Language

type Section struct {
	ID           int           `db:"id" json:"id"`
	LanguageCode string        `db:"language_code"`
	ContentType  string        `db:"content_type" json:"content_type"`
	Title        string        `db:"title" json:"title"`
	Position     int           `db:"position" json:"position"`
	DefaultBlock *DefaultBlock `json:"default_block,omitempty"`
	// TODO: other content types
}

type SectionList []*Section

type DefaultBlock struct {
	ID          int    `db:"id" json:"id"`
	SectionID   int    `db:"section_id" json:"section_id"`
	Address     string `db:"address" json:"address"`
	Timetable   string `db:"timetable" json:"timetable"`
	Phone       string `db:"phone" json:"phone"`
	Description string `db:"description" json:"description"`
}
