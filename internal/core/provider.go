package core

import "database/sql"

// Проводник (вообще, конечно, Поставщик, но для избежания путаницы с Поставщиками в энергетике назван по другому)
type Provider struct {
	LINK             int
	Name             string
	ConnectionString sql.NullString
	Definition       string

	Type     string
	TypeLink int

	ProviderType
}

type ProviderType struct {
	LINK  int
	Name  string
	Const string
}
