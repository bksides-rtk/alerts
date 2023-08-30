package main

import "github.com/RTK-Tickets/alerts/server"

func main() {
	srv := server.NewServer(server.ServerConfig{
		Addr: ":8080",
		Router: server.RouterConfig{
			GraphQLEndpointPath: "/graphql",
		},
	})

	srv.ListenAndServe()
}
