package model

type Language struct {
	Code     string      `db:"code" json:"code"`
	Sections SectionList `json:"sections"`
}

type LanguageList []*Language

type Section struct {
	ID           int         `db:"id" json:"-"`
	LanguageCode string      `db:"language_code" json:"-"`
	ContentType  string      `db:"content_type" json:"content_type"`
	Title        string      `db:"title" json:"title"`
	Icon         string      `db:"icon" json:"icon"`
	Position     int         `db:"position" json:"-"`
	Data         interface{} `json:"data,omitempty"`
}

type SectionList []*Section
