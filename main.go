package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
	"strconv"
)
type data struct{
	Visits string `json:"visits"`
}
func main(){
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