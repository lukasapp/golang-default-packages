package main

import (
	"fmt"

	"github.com/lukasapp/golib/middleware"
)

func main() {
	corsOptions := &middleware.CorsOptions{
		Origin: "*", Methods: "OPTIONS,POST,GET,PUT,DELETE", Headers: "*",
	}

	middlewareStack := middleware.CreateStack(
		middleware.Cors(corsOptions),
		middleware.Logging,
	)

	fmt.Println("Middleware stack", middlewareStack)
}
