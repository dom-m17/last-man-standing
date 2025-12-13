package logger

import "go.uber.org/zap"

type ServiceInterface interface {
	Error(args ...any)
	Warn(args ...any)
	Info(args ...any)
}

type Service struct {
	logger *zap.SugaredLogger
}

func NewService() *Service {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	return &Service{
		logger: sugar,
	}
}
