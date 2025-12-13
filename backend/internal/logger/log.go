package logger

func (s *Service) Error(args ...any) {
	s.logger.Error(args...)
}

func (s *Service) Warn(args ...any) {
	s.logger.Warn(args...)
}

func (s *Service) Info(args ...any) {
	s.logger.Info(args...)
}
