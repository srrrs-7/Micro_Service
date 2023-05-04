package handle

func (s *server) AuthRouter() {
	s.router.GET("/auth", auth.NewAuth())
}
