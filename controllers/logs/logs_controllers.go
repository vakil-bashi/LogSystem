package logs

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vakil-bashi/log-system/logger"
	"github.com/vakil-bashi/log-system/services"
	"github.com/vakil-bashi/log-system/utils/responses"
	"net/http"
)

func InsertLogs(c *gin.Context) {
	logger.Info("Enter to InsertLogs controller successfully")

	apiKey := c.GetHeader("X-API-Key")

	if apiKey == "" || apiKey != "wzw0ReNLnrOnPSvpG4i1UT7hKTitSYnb3F8BjM0mS6UU7YH6QjqczD3MJhsp7HI44VmuETG9b1Na0S3LZCFtRtvFvt2vvGajLYRCqNShm7nluZD62o4LxmBaRZHukLLL" {
		logger.Error("request is Unauthorized", errors.New("api-key is not valid"))
		restError := responses.NewBadRequestError("API-Key is not Valid", "please makes sure to send correct api-key.", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, restError)
		c.Abort()
		return
	}

	logREQ := make(map[string]interface{})
	if err := c.ShouldBindJSON(&logREQ); err != nil {
		logger.Error("error when trying to bind json", err)
		restError := responses.NewBadRequestError("Invalid json structure", "please make sure to send correct json body", http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, restError)
		return
	}

	serviceErr := services.LogsService.InsertES(logREQ)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("logs Inserted into elastic-search successfully", "", nil)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.JSON(http.StatusOK, ok)
	logger.Info("Close from InsertLogs controller successfully")
}
