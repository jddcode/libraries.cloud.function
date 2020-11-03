package cloudFunction

import (
	"errors"
	"net/http"
	"os"
	"strings"
)

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

		handler, params, err := s.findHandler(r)

		if err != nil {

			w.WriteHeader(404)
			w.Write([]byte("No provider could be found"))
			return
		}

		req := Request{ input: r, output: w, args: params }
		response := handler(req)

		w.WriteHeader(response.StatusCode)

		if len(os.Getenv("BuildDate")) > 0 {

			w.Header().Set("X-Build-Date", os.Getenv("BuildDate"))
		}

		s.corsInjector(w)
		w.Write(response.Content)
	}

	func (s *Service) findHandler(r *http.Request) (Handler, map[string]string, error) {

		for _, route := range s.routes {

			match, args := s.isAMatch(route.Path, r.URL.Path)

			if match {

				return route.Handler, args, nil
			}
		}

		return nil, nil, errors.New("no_handler")
	}

	func (s *Service) isAMatch(path, url string) (bool, map[string]string) {

		urlBits := strings.Split(strings.ToLower(strings.Trim(url, "/")), "/")
		pathBits := strings.Split(strings.ToLower(strings.Trim(path, "/")), "/")

		if len(urlBits) != len(pathBits) {

			return false, nil
		}

		match := true
		args := make(map[string]string)
		for pos, bit := range pathBits {

			if match && len(bit) > 1 && bit[0:1] == ":" {

				args[bit[1:]] = urlBits[pos]

			} else {

				if urlBits[pos] != pathBits[pos] {

					args = nil
					match = false
				}
			}
		}

		return match, args
	}

	func (s *Service) corsInjector(w http.ResponseWriter) {

		w.Header().Set("Access-Control-Allow-Headers","*")
		w.Header().Set("Access-Control-Allow-Origin","*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	}

	type route struct {

		Path string
		Handler Handler
	}

	type Handler func(request Request) Response
