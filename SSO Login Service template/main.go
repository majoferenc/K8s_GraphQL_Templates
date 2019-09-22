package main

import (
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"log"
	"net/http"
)

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			fetchQuery(),
		),
		Mutation: rootMutation(
			createEntityType("entityResult"),
		),
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
		GraphiQL:true,
	})
	http.Handle("/graphql", CorsMiddleware(handler))
	log.Println("GraphQL API started at http://localhost:8091/graphql")
	log.Fatal(http.ListenAndServe(":8091", nil))
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w,r)
	})
}

type EntityMetadata struct {
	Id          int `json:"id"`
	Name string `json:"name"`
}

func fetchQuery() graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "QueryType", Fields: graphql.Fields{
		"fetchById": &graphql.Field{
			Type: getEntityTypeById("entityMetadata"),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"]
				v, _ := id.(int)
				name := params.Args["name"].(string)
				log.Printf("fetching entity metadata with id: %d", v)
				return fetchEntity(v,name)
			},
		},
		"fetchAll": &graphql.Field{
			Type: getEntities("entityArray"),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				log.Print("fetching all entities")
				return fetchEntities()
			},
		},
	}}
}

func rootMutation(entityType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createEntity": &graphql.Field{
				Type: entityType, // the return type for this field
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					// marshall and cast the argument value
					id, _ := params.Args["id"].(int)
					name, _ := params.Args["name"].(string)

					// perform mutation operation here
					// for e.g. create a Todo and save to DB.
					newEntity := &EntityMetadata{
						Id: id,
						Name: name,
					}

					// return the new Todo object that we supposedly save to DB
					// Note here that
					// - we are returning a `Todo` struct instance here
					// - we previously specified the return Type to be `todoType`
					// - `Todo` struct maps to `todoType`, as defined in `todoType` ObjectConfig`
					return newEntity, nil
				},
			},
		},
	})
}

func getEntities(name string) *graphql.List {
	entityType := graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Description: "Returning all entities",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	return graphql.NewList(entityType)
}

func getEntityTypeById(name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Description: "Returning entity by id",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}

func createEntityType(name string) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: name,
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})
}


func fetchEntity(id int,name string) (*EntityMetadata, error) {
	result := EntityMetadata{Id: id,Name: name}
	return &result, nil
}

func fetchEntities() (*[]EntityMetadata, error) {
	arr := make([]EntityMetadata,0)
	entity1 := EntityMetadata{Id: 0,Name: "namee"}
	entity2 := EntityMetadata{Id: 1,Name: "text"}
	arr = append(arr, entity1)
	arr = append(arr, entity2)
	return &arr, nil
}