package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

type Info struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, &Info{
			Name: "kanna",
			Age:  24,
		})
	})

	r.Run("localhost:9090")
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	data, _ := json.Marshal(&Info{
//		Name: "kanna",
//		Age:  24,
//	})
//	w.Write(data)
//}
//
//func main() {
//	http.HandleFunc("/", handler)
//	http.ListenAndServe("localhost:8080", nil)
//
//}
