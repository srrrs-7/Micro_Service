package csv

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCsvHandler(ctx *gin.Context) {
	var req csvRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// csv column length check headers == values[i]
	for i := 0; i < len(req.Headers); i++ {
		if len(req.Values[i]) != len(req.Headers) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid csv column length"})
			return
		}
	}

	createCsv(req.Headers, req.Values)

	ctx.JSON(http.StatusOK, gin.H{"message": "success create csv file"})
}
