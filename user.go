package main

import (
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
)

var userType *graphql.Object
var usersType *graphql.Object

var nodeDefinitions *relay.NodeDefinitions

var Schema graphql.Schema

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "A person who uses our app",
	Fields: graphql.Fields{
		"id": relay.GlobalIDField("User", nil),
		"email": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Email, nil
				}
				return nil, nil
			},
		},
		"surname": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Surname, nil
				}
				return nil, nil
			},
		},
		"lastname": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if user, ok := p.Source.(*User); ok == true {
					return user.Lastname, nil
				}
				return nil, nil
			},
		},
	},
})

func init() {

	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ct context.Context) (interface{}, error) {
			resolvedID := relay.FromGlobalID(id)
			if resolvedID.Type == "User" {
				i, err := strconv.Atoi(resolvedID.ID)
				if err != nil {
					return nil, err
				}
				return GetUserByID(i)
			}
			return nil, nil
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			switch p.Value.(type) {
			case *User:
				return userType
			}
			return nil
		},
	})

	usersType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Users",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: graphql.NewList(UserType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// add business logic to retrieve list of Posts
					return GetUsers()
				},
			},
		},
	})

	/**
	 * This is the type that will be the root of our query,
	 * and the entry point into our schema.
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": nodeDefinitions.NodeField,

			// Add you own root fields here
			"viewer": &graphql.Field{
				Type: usersType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetUsers()
				},
			},
		},
	})

	/**
	* Finally, we construct our schema (whose starting query type is the query
	* type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
	if err != nil {
		panic(err)
	}
}
