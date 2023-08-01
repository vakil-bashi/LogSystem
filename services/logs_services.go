package services

import (
	"github.com/vakil-bashi/log-system/logger"
	"github.com/vakil-bashi/log-system/utils/responses"
)

var (
	LogsService logsServiceInterface = &logsService{}
)

type logsService struct{}

type logsServiceInterface interface {
}

func (s *logsService) InsertES() *responses.Response {
	logger.Info("Enter to InsertES service successfully")
	logger.Info("Close from InsertES service successfully")
	return nil
}
