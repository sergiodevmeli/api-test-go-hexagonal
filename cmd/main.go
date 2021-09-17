package main

import (
	"api_test_hexagonal/pkg"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var _ = godotenv.Load(".env")

func main()  {
	mux := http.NewServeMux()
	pkg.IndexRouter(mux)
	port := fmt.Sprintf(":%s", os.Getenv("port_server"))
	log.Fatal(http.ListenAndServe(port, mux))
}

