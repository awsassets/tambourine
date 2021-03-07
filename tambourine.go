package tambourine

// Service will be the main service
// loop for Tambourine
type Service struct {
	logger  Logger
	options *ServiceOptions
}

type ServiceOptions struct {
	//
}

func New(opt *ServiceOptions) *Service {
	return &Service{
		options: opt,
		logger:  &StdOutLogger{},
	}
}

func (s *Service) Run() error {
	s.logger.Info("starting service...")
	for {
		// TODO Service loop
	}
	return nil
}
