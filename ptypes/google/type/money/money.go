package money

import (
	"github.com/graphql-go/graphql"
	"google.golang.org/genproto/googleapis/type/money"
)

// Expose Google defined ptypes as this package types
type Money = money.Money

var (
	gql__type_Money  *graphql.Object
	gql__input_Money *graphql.InputObject
)

func Gql__type_Money() *graphql.Object {
	if gql__type_Money == nil {
		gql__type_Money = graphql.NewObject(graphql.ObjectConfig{
			Name: "Google_type_Money",
			Fields: graphql.Fields{
				"currencyCode": &graphql.Field{
					Type:        graphql.String,
					Description: `The three-letter currency code defined in ISO 4217.`,
				},
				"units": &graphql.Field{
					Type: graphql.Int,
					Description: `The whole units of the amount.
For example if ` + "`" + `currencyCode` + "`" + ` is ` + "`" + `"USD"` + "`" + `, then 1 unit is one US dollar.`,
				},
				"nanos": &graphql.Field{
					Type: graphql.Int,
					Description: `Number of nano (10^-9) units of the amount.
The value must be between -999,999,999 and +999,999,999 inclusive.
If ` + "`" + `units` + "`" + ` is positive, ` + "`" + `nanos` + "`" + ` must be positive or zero.
If ` + "`" + `units` + "`" + ` is zero, ` + "`" + `nanos` + "`" + ` can be positive, zero, or negative.
If ` + "`" + `units` + "`" + ` is negative, ` + "`" + `nanos` + "`" + ` must be negative or zero.
For example $-1.75 is represented as ` + "`" + `units` + "`" + `=-1 and ` + "`" + `nanos` + "`" + `=-750,000,000.`,
				},
			},
			Description: `Represents an amount of money with its currency type.`,
		})
	}
	return gql__type_Money
}

func Gql__input_Money() *graphql.InputObject {
	if gql__input_Money == nil {
		gql__input_Money = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Google_input_Money",
			Fields: graphql.InputObjectConfigFieldMap{
				"currencyCode": &graphql.InputObjectFieldConfig{
					Type:        graphql.String,
					Description: `The three-letter currency code defined in ISO 4217.`,
				},
				"units": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
					Description: `The whole units of the amount.
For example if ` + "`" + `currencyCode` + "`" + ` is ` + "`" + `"USD"` + "`" + `, then 1 unit is one US dollar.`,
				},
				"nanos": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
					Description: `Number of nano (10^-9) units of the amount.
The value must be between -999,999,999 and +999,999,999 inclusive.
If ` + "`" + `units` + "`" + ` is positive, ` + "`" + `nanos` + "`" + ` must be positive or zero.
If ` + "`" + `units` + "`" + ` is zero, ` + "`" + `nanos` + "`" + ` can be positive, zero, or negative.
If ` + "`" + `units` + "`" + ` is negative, ` + "`" + `nanos` + "`" + ` must be negative or zero.
For example $-1.75 is represented as ` + "`" + `units` + "`" + `=-1 and ` + "`" + `nanos` + "`" + `=-750,000,000.`,
				},
			},
			Description: `Represents an amount of money with its currency type.`,
		})
	}
	return gql__input_Money
}
