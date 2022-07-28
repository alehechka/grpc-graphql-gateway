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
					Type: graphql.String,
				},
				"units": &graphql.Field{
					Type: graphql.Int,
				},
				"nanos": &graphql.Field{
					Type: graphql.Int,
				},
			},
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
					Type: graphql.String,
				},
				"units": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"nanos": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
			},
		})
	}
	return gql__input_Money
}
