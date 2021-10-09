package main

import(
  "fmt"
  "log"
  "context"
  "time"
  "net/http"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
    id string
    caption string
    imageURL string
    postedTimestamp string
}

type post struct {
    id string
    name string
    email string
    password string
}

type returnType interface{}
func intialiseDatabase (returnType, returnType) {

    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Aniket:Shubhamjazz1@instagram.56kvi.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
    if err!=nil{log.Fatal(err)} // IN CASE OF ERROR

    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    // CONNECTING WITH THE DATABASE

    err = client.Connect(ctx)
    if err!=nil{log.Fatal(err)} // INCASE OF ERROR
    defer client.Disconnect(ctx)  // RUNS IN THE END, DISCONNECTS THE CONNECTION
    instagram_database := client.Database("instagram")

    // CONNECT WITH A COLLECTION

    usersTable := instagram_database.Collection("users_collection")
    return usersTable,instagram_database

    return

}

func InserUserData(w http.ResponseWriter, r *http.Request){
    if r.Method == "POST" {
        r.ParseForm() // PARSES CLIENT SIDE INPUT
        users_table,instagram_database = intialiseDatabase();
        is_duplicate = instagram_database.system.indexes.find({"id":"r.Form["id"]", "ns":"instagram_database.users_table"}).count()

        if is_dupilcate==0{
            newUser := user{r.Form["id"], r.Form["name"], r.Form["email"], r.Form["password"]}

            // inserting into the database
            result, err := users.InsertOne(ctx, newUser)
            if err!=nil { log.Fatal(err) }

            fmt.FprintF(w,"Account Added")
        }else{
            fmt.FprintF(w,"This User Id already exists")
        }
    }
}

func GetUser (w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        r.ParseForm()
        users_table,instagram_database = intialiseDatabase();
        is_duplicate = instagram_database.system.indexes.find({"id":"r.Form["id"]", "ns":"instagram_database.users_table"}).count()

        result:= client.Database(instagram_database).Collection("users_table").FindOne(context.Background(), bson.M{"_id": r.Form["id"]})
        user := model.User{}
        result.Decode(user)
    }
}

func (u user) PostImage (w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "uploading this file")

    r.ParseMultipartForm(10 << 20)
    file, handler, err := r.FormFile("imagePost")
    if err != nil {
        fmt.Println("error")
        fmt.Println(err)
    }
    defer file.Close()
    fmt.Printf("The image has been posted : %+v", handler.Filename)

    imageTemp, err := ioutil.TempFile("temp-images","upload-*.png")
    if err != nil {
        fmt.Println(err)
        return
    }defer tempFile.Close()

    bytesOfFile, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }




    imageTemp.Write(bytesOfFile)

    fmt.Fprintf(w, "Uplaoded")
}

func endpoints (int) {
    //USER RELATED
    http.HandleFunc("/users",InsertUserData)
    http.HandleFunc("/users/<users_id>", GetUser)

    // POST RELATED
    http.HandleFunc("/post", PostImage)
    http.HandleFunc("/posts/<posts_id>", GetPostId)
    http.HandleFunc("/posts/users/<id>", UserPostList)
}

func main(){
    endpoints()
}
