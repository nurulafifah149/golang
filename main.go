package main

import "github.com/nurulafifah149/golang/server"

// @host      localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization

func main() {
	server.Serve()
}
