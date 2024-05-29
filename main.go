package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	// Additional dependencies
	_ "github.com/aws/aws-sdk-go"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator"
	_ "github.com/googleapis/gax-go/v2"
	_ "github.com/gorilla/websocket"
	_ "github.com/grpc-ecosystem/go-grpc-middleware"
	_ "github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv"
	_ "github.com/prometheus/client_golang/prometheus"
	_ "github.com/sirupsen/logrus"
	_ "github.com/spf13/viper"
	_ "golang.org/x/image/draw"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
