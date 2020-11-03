package cloudFunction

import "net/http"

	func Serve(handler Handler, w http.ResponseWriter, r *http.Request) {

		request := Request{ input: r, output: w }
		response := handler(request)

		if response.StatusCode < 100 {

			response.StatusCode = http.StatusBadRequest
			response.Content = []byte("Handler failure")
		}

		w.WriteHeader(response.StatusCode)
		w.Write(response.Content)
	}

	type Handler func(request Request) Response
