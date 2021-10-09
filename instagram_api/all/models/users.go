package main

import (
    "fmt"
    "log"
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

)

func main() {

    // GETTING THE CLIENT WHICH HELPS IN CONNECTION

    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Aniket:Shubhamjazz1@instagram.56kvi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
    if err!=nil{log.Fatal(err)} // IN CASE OF ERROR

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    // CONNECTING WITH THE DATABASE

    err = client.Connect(ctx)
    if err!=nil{log.Fatal(err)} // INCASE OF ERROR
    defer client.Disconnect(ctx)  // RUNS IN THE END, DISCONNECTS THE CONNECTION
    instagram_database := client.Database("instagram")

    // CONNECT WITH A COLLECTION

    users := instagram_database.Collection("users_collection")

    // INSERT USER DATA
/*
    result, err := users.InsertOne(ctx, bson.D{{"keyOne", "valueOne"},{"keyTwo", "valueTwo"}})
    if err!=nil { log.Fatal(err) }
    id := result.InsertedID*/
    // RETURN TRUE

}
