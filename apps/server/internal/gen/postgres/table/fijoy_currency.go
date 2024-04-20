//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var FijoyCurrency = newFijoyCurrencyTable("public", "fijoy_currency", "")

type fijoyCurrencyTable struct {
	postgres.Table

	// Columns
	ID   postgres.ColumnString
	Code postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FijoyCurrencyTable struct {
	fijoyCurrencyTable

	EXCLUDED fijoyCurrencyTable
}

// AS creates new FijoyCurrencyTable with assigned alias
func (a FijoyCurrencyTable) AS(alias string) *FijoyCurrencyTable {
	return newFijoyCurrencyTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FijoyCurrencyTable with assigned schema name
func (a FijoyCurrencyTable) FromSchema(schemaName string) *FijoyCurrencyTable {
	return newFijoyCurrencyTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FijoyCurrencyTable with assigned table prefix
func (a FijoyCurrencyTable) WithPrefix(prefix string) *FijoyCurrencyTable {
	return newFijoyCurrencyTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FijoyCurrencyTable with assigned table suffix
func (a FijoyCurrencyTable) WithSuffix(suffix string) *FijoyCurrencyTable {
	return newFijoyCurrencyTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFijoyCurrencyTable(schemaName, tableName, alias string) *FijoyCurrencyTable {
	return &FijoyCurrencyTable{
		fijoyCurrencyTable: newFijoyCurrencyTableImpl(schemaName, tableName, alias),
		EXCLUDED:           newFijoyCurrencyTableImpl("", "excluded", ""),
	}
}

func newFijoyCurrencyTableImpl(schemaName, tableName, alias string) fijoyCurrencyTable {
	var (
		IDColumn       = postgres.StringColumn("id")
		CodeColumn     = postgres.StringColumn("code")
		allColumns     = postgres.ColumnList{IDColumn, CodeColumn}
		mutableColumns = postgres.ColumnList{CodeColumn}
	)

	return fijoyCurrencyTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:   IDColumn,
		Code: CodeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}