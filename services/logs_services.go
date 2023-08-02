package services

import (
	"encoding/json"
	"github.com/vakil-bashi/log-system/domains/requests"
	"github.com/vakil-bashi/log-system/logger"
	"github.com/vakil-bashi/log-system/utils/responses"
	"net/http"
)

var (
	LogsService logsServiceInterface = &logsService{}
)

type logsService struct{}

type logsServiceInterface interface {
	InsertES(map[string]interface{}) *responses.Response
}

func (s *logsService) InsertES(logREQ map[string]interface{}) *responses.Response {
	logger.Info("Enter to InsertES service successfully")

	req := requests.Req{
		URL:    "http://localhost:9200/logs/_doc/",
		Method: "POST",
	}

	simpleAuthentication := true
	body, err := req.POST(logREQ, nil, &simpleAuthentication)
	if err != nil {
		return err
	}

	ret := make(map[string]interface{})
	if marshalErr := json.Unmarshal(body, &ret); marshalErr != nil {
		logger.Error("error when trying to marshal response to struct", marshalErr)
		err := responses.NewBadRequestError("could not marshal the response struct", "bad-request", http.StatusBadRequest)
		return err
	}
	logger.Info("Close from InsertES service successfully")
	return nil
}
