package core

import "database/sql"

type TemplateGroup struct {
	LINK int
	Name string
}

// Шаблон перивичного импорта
type Template struct {
	Group TemplateGroup

	LINK       int
	Name       string
	SystemName string
	Order      int
	FromQuery  string
	ToQuery    string

	GroupLINK int

	Rows []TemplateRow
}

// Строка шаблона первичного импорта
type TemplateRow struct {
	LINK             int
	IsMandatory      bool   `json:"mandatory"`
	Order            int    `json:"order"`
	Name             string `json:"name"`
	SystemName       string `json:"systemname"`
	Definition       sql.NullString
	Comment          string        `json:"comment"`
	SQLType          string        `json:"type"`
	Precise          sql.NullInt32 // До запятой
	Precise2         sql.NullInt32 // После запятой
	TemplateBaseLINK int           `json:"templatebase"`
	TemplateRowLINK  int           `json:"templaterow"`
	DictionaryLINK   int           `json:"dictionary"`

	Template
	Dictionary
}
