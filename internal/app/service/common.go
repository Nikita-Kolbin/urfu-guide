package service

func (s *Service) GetJWTSecret() string {
	return s.jwtSecret
}
