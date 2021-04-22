package server

import (
	"net/http"
	"time"
)

type option func(*Server)

func WithTimeout(d time.Duration) option {
	return func(s *Server) {
		s.WithTimeout(d)
	}
}

func WithAddr(addr string) option {
	return func(s *Server) {
		if addr == "" {
			addr = ":4242"
		}

		s.WithAddr(addr)
	}
}

type Server struct {
	*http.Server
}

func NewServer(options ...option) *Server {
	srv := &Server{&http.Server{}}

	for _, opt := range options {
		opt(srv)
	}

	return srv
}

func (s *Server) WithTimeout(d time.Duration) {
	s.ReadTimeout = time.Second * d
	s.WriteTimeout = time.Second * d
}

func (s *Server) WithAddr(addr string) {
	s.Addr = addr
}
