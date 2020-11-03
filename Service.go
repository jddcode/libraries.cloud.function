package cloudFunction

import "net/http"

	type Service struct {


	}

	func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("works"))
	}
