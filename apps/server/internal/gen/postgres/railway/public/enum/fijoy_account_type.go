//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var FijoyAccountType = &struct {
	Chequing postgres.StringExpression
	Savings  postgres.StringExpression
	Credit   postgres.StringExpression
	Mortgage postgres.StringExpression
	Gic      postgres.StringExpression
	Tfsa     postgres.StringExpression
	Rrsp     postgres.StringExpression
	Fhsa     postgres.StringExpression
	Cash     postgres.StringExpression
}{
	Chequing: postgres.NewEnumValue("chequing"),
	Savings:  postgres.NewEnumValue("savings"),
	Credit:   postgres.NewEnumValue("credit"),
	Mortgage: postgres.NewEnumValue("mortgage"),
	Gic:      postgres.NewEnumValue("gic"),
	Tfsa:     postgres.NewEnumValue("tfsa"),
	Rrsp:     postgres.NewEnumValue("rrsp"),
	Fhsa:     postgres.NewEnumValue("fhsa"),
	Cash:     postgres.NewEnumValue("cash"),
}