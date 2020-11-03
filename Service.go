package cloudFunction

import "net/http"

	type Service struct {

		routes []route
	}

	func (s *Service) AddRoute(path string, handler Handler) {

		if len(s.routes) < 1 {

			s.routes = make([]route, 0)
		}

		s.routes = append(s.routes, route{ Path: path, Handler: handler })
	}

	func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {

		response := s.routes[0].Handler(Request{ input: r, output: w })
		w.WriteHeader(response.StatusCode)
		w.Write(response.Content)
	}

	type route struct {

		Path string
		Handler Handler
	}
