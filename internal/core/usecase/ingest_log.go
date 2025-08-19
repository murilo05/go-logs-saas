package usecase

import "go.uber.org/zap"

type LogService struct {
	logger *zap.SugaredLogger
}

func NewLogService(logger *zap.SugaredLogger) *LogService {
	return &LogService{
		logger: logger,
	}
}
