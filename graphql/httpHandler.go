package graphql

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type GraphQLHandler struct {
	Schema *graphql.Schema
}

type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func (h *GraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var graphQlRequest GraphQLRequest
	err := json.NewDecoder(r.Body).Decode(&graphQlRequest)
	if err != nil {
		http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:         *h.Schema,
		RequestString:  graphQlRequest.Query,
		VariableValues: graphQlRequest.Variables,
		OperationName:  graphQlRequest.OperationName,
	})

	if len(result.Errors) > 0 {
		for _, err := range result.Errors {
			fmt.Println(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
