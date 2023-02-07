package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khalil-farashiani/products-service/internals/infrastructure/persistence/mongodb"
	"github.com/khalil-farashiani/products-service/internals/usecase/product"
	"github.com/khalil-farashiani/products-service/router"
	"log"
	"os"
)

const localHost = "localhost:8080"

func main() {

	mongoDsn := os.Getenv("mongoDsn")
	store, err := mongodb.NewProductRepository(mongoDsn)
	if err != nil {
		log.Fatalln("error in connecting to Db")
	}

	product.NewProductUseCase(store)

	r := gin.Default()
	router.GetRouts(r)

	r.Run(localHost)
}
