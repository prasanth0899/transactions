package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"transactions/services"
)

func GetTransactions(c *gin.Context) {
	params := c.Params
	orgId, _ := params.Get("orgId")
	limitString, _ := params.Get("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Printf("limit should be in integer")
	}
	transactions := services.GetTransactions(orgId, limit, nil)
	c.JSON(http.StatusOK, transactions)
}
