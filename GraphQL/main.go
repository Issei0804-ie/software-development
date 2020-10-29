package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"string"`
	Pass string `json:"password"`
}

var users = []user{
	{Id:   1, Name: "hoge", Pass: "08-16"},
	{Id: 2, Name: "fuga", Pass: "06-17"},
	{Id: 3, Name: "piyo", Pass: "05-22"},
	{Id: 4, Name: "hogehoge", Pass: "07-23"},
}

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"pass": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/* Get (read) single user by id
			   http://localhost:8080/user?query={user(id:1){name,pass}}
			   $ curl "http://localhost:8080/user?query=%7Buser%28id:1%29%7Bname,pass%7D%7D"
			*/
			"user": &graphql.Field{
				Type:        productType,
				Description: "Get user by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// Find product
						return users[id], nil
					}
					return nil, nil
				},
			},
			/* Get (read) product list
			   http://localhost:8080/user?query={list{id,name,pass}}
			*/
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get product list",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return users, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
