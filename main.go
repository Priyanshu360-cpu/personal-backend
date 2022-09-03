package main
import (
	"context"
	"fmt"
	"log"
	"time"
    "net/http"
    "github.com/gin-gonic/gin"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
type data struct{
	Visits string `json:"visits"`
}

type reviews struct{
	Message string `json:"message"`
	Likes int `json:"likes"`
}

func main(){
	client, err := mongo.NewClient(options.Client().ApplyURI(""))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())

if err != nil {
    log.Fatal(err)
}
databases, err := client.ListDatabaseNames(ctx, bson.M{})
if err != nil {
    log.Fatal(err)
}
fmt.Println(databases)
	var a = data{Visits: "0"}
	r := gin.Default()
	r.GET("/visits", func(c *gin.Context){
		e,err:= strconv.Atoi(a.Visits)
		if err == nil {
			e++
		}
		a.Visits = strconv.Itoa(e)
		c.JSON(http.StatusOK, a)
	})
   r.Run("localhost:8080")
}