package jsonHandler

import (
	"fmt"
	"net/http"
	dataStruct "server/dataStruct"
	postgresDB "server/postgresDB"

	"github.com/gin-gonic/gin"
)

func ReceiveJSON() {
	var jsonData dataStruct.SystemInfo

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.POST("/rpi", func(c *gin.Context) {
		if err := c.BindJSON(&jsonData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {

			fmt.Printf("Received Info: %+v\n", jsonData)
			c.JSON(http.StatusOK, gin.H{"message": "JSON data received successfully"})
			postgresDB.CheckPrimaryKey(jsonData)
			postgresDB.CloseDB()
		}
	})

	r.Run(":8080")
}
