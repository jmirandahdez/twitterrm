package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarBD()
var cleintOptions = options.Client().ApplyURI("mongodb+srv://sa:Abril190@twitterrm.bhjolmp.mongodb.net/?retryWrites=true&w=majority")

//ConectarBD Conexiona a la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), cleintOptions)
	if err != nil {
		log.Println("Error")

		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println("Error1")

		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa")
	return client
}

// CheckConnection valida la conexion
// Return 0 en caso de error, 1 exitoso
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
