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

var FijoyWorkspace = newFijoyWorkspaceTable("public", "fijoy_workspace", "")

type fijoyWorkspaceTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	Name      postgres.ColumnString
	Namespace postgres.ColumnString
	CreatedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FijoyWorkspaceTable struct {
	fijoyWorkspaceTable

	EXCLUDED fijoyWorkspaceTable
}

// AS creates new FijoyWorkspaceTable with assigned alias
func (a FijoyWorkspaceTable) AS(alias string) *FijoyWorkspaceTable {
	return newFijoyWorkspaceTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FijoyWorkspaceTable with assigned schema name
func (a FijoyWorkspaceTable) FromSchema(schemaName string) *FijoyWorkspaceTable {
	return newFijoyWorkspaceTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FijoyWorkspaceTable with assigned table prefix
func (a FijoyWorkspaceTable) WithPrefix(prefix string) *FijoyWorkspaceTable {
	return newFijoyWorkspaceTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FijoyWorkspaceTable with assigned table suffix
func (a FijoyWorkspaceTable) WithSuffix(suffix string) *FijoyWorkspaceTable {
	return newFijoyWorkspaceTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFijoyWorkspaceTable(schemaName, tableName, alias string) *FijoyWorkspaceTable {
	return &FijoyWorkspaceTable{
		fijoyWorkspaceTable: newFijoyWorkspaceTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newFijoyWorkspaceTableImpl("", "excluded", ""),
	}
}

func newFijoyWorkspaceTableImpl(schemaName, tableName, alias string) fijoyWorkspaceTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		NameColumn      = postgres.StringColumn("name")
		NamespaceColumn = postgres.StringColumn("namespace")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		allColumns      = postgres.ColumnList{IDColumn, NameColumn, NamespaceColumn, CreatedAtColumn}
		mutableColumns  = postgres.ColumnList{NameColumn, NamespaceColumn, CreatedAtColumn}
	)

	return fijoyWorkspaceTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Name:      NameColumn,
		Namespace: NamespaceColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}