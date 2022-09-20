package bd

import (
	"context"
	"fmt"

	"time"

	"github.com/jmirandahdez/twitterrm/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscoPerfil(Id string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("twitterrm")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(Id)

	condicion := bson.M{"ID": objID}
	err := col.FindOne(ctx, condicion).Decode(&perfil)

	if err == nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, err
	}

	perfil.Password = ""
	return perfil, nil
}
