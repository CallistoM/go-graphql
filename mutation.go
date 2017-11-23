package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type: UserType,
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Description: "New User Email",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"surname": &graphql.ArgumentConfig{
					Description: "New User Surname",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"lastname": &graphql.ArgumentConfig{
					Description: "New User Lastname",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email := p.Args["email"].(string)
				surname := p.Args["surname"].(string)
				lastname := p.Args["lastname"].(string)
				user := &User{
					Email:    email,
					Surname:  surname,
					Lastname: lastname,
				}
				err := InsertUser(user)
				return user, err
			},
		},
		"removeUser": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "User ID to remove",
					Type:        graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				i := p.Args["id"].(string)
				id, err := strconv.Atoi(i)
				if err != nil {
					return nil, err
				}
				err = RemoveUserByID(id)
				return (err == nil), err
			},
		},
	},
})
