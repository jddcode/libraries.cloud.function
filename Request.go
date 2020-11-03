package cloudFunction

import (
	"io/ioutil"
	"net/http"
)

	type Request struct {

		input *http.Request
		output http.ResponseWriter
		params map[string]string
		body []byte
	}

	func (r *Request) Error(msg string) Response {

		return Response{ StatusCode: http.StatusBadRequest, Content: []byte(msg) }
	}

	func (r *Request) Success() Response {

		return Response{ StatusCode: http.StatusOK }
	}

	func (r *Request) SuccessWithMsg(msg string) Response {

		return Response{ StatusCode: http.StatusOK, Content: []byte(msg) }
	}

	func (r *Request) SuccessWithBytes(content []byte) Response {

		return Response{ StatusCode: http.StatusOK, Content: content }
	}

	func (r *Request) Body() ([]byte, error) {

		if len(r.body) > 0 {

			return r.body, nil
		}

		bytes, err := ioutil.ReadAll(r.input.Body)

		if err != nil {

			return nil, err
		}

		r.body = bytes
		return bytes, nil
	}

	func (r *Request) HasBody() bool {

		body, err := r.Body()

		if err != nil || len(body) < 1 {

			return false
		}

		return true
	}