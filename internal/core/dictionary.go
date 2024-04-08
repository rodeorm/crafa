package core

import (
	"database/sql"
	"encoding/json"

	"github.com/pkg/errors"
)

type Dictionary struct {
	LINK        int
	Number      int
	Name        string
	SystemName  string
	Definition  string
	CodeColumn  sql.NullString
	ValueColumn sql.NullString

	TypeLINK     int
	TypeName     string
	ProviderLINK int
	ProviderName string
	JRows        sql.NullString
	SyncTime     sql.NullTime

	DictionaryType
	DictionaryGroup
	Provider
	Rows []DictionaryRow
	History
}

type DictionaryType struct {
	LINK int
	Name string
}

type DictionaryGroup struct {
	LINK int
	Name string
}

type DictionaryRow struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

func (d *Dictionary) GetRowsFromJSON() error {
	DictionaryRows := make([]DictionaryRow, 0)
	err := json.Unmarshal([]byte(d.JRows.String), &DictionaryRows)
	if err != nil {
		return errors.Wrap(err, "ошибка при попытке получения строк справочника из json")
	}
	d.Rows = DictionaryRows
	return nil
}
