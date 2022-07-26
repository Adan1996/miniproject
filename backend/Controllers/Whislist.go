package Controllers

import (
	"fmt"
	"net/http"

	"example.com/m/Models"
	"github.com/gin-gonic/gin"
)

func GetWhislist(c *gin.Context) {
	var whistlist []Models.Whislist
	err := Models.GetAllWhislist(&whistlist)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, whistlist)
	}
}
func CreateWhistlist(c *gin.Context) {
	var whislist Models.Whislist
	c.BindJSON(&whislist)
	err := Models.CreateWhislist(&whislist)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, whislist)
	}
}
