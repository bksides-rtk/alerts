package graphql

import (
	"github.com/RTK-Tickets/alerts/models"
	"github.com/graphql-go/graphql"
)

var Schema graphql.Schema

func init() {
	var AlertType = graphql.NewObject(graphql.ObjectConfig{
		Name:   "Alert",
		Fields: graphql.BindFields(models.Alert{}),
	})

	var AlertSearchType = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "AlertSearch",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"severity": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"loud": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"claimedBy": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"source": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"dismissed": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	})

	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"alerts": &graphql.Field{
					Name: "alerts",
					Args: graphql.FieldConfigArgument{
						"match": &graphql.ArgumentConfig{
							Type:        AlertSearchType,
							Description: "Alert against which to match",
						},
					},
					Type:        graphql.NewList(AlertType),
					Description: "Get a list of alerts matching the given alert.",
					Resolve:     getAlerts,
				},
				"alert": &graphql.Field{
					Name: "alert",
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type:        graphql.Int,
							Description: "ID of the alert to retrieve",
						},
					},
					Type:        AlertType,
					Description: "Get an alert by ID.",
					Resolve:     getAlert,
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"createAlert": &graphql.Field{
					Name: "createAlert",
					Args: graphql.FieldConfigArgument{
						"alert": &graphql.ArgumentConfig{
							Type:        AlertSearchType,
							Description: "Alert to create",
						},
					},
					Type:        AlertType,
					Description: "Create a new alert.",
					Resolve:     createAlert,
				},
			},
		}),
	})

	if err != nil {
		panic(err)
	}
}

var AlertQueryHandler = &GraphQLHandler{
	Schema: &Schema,
}
