package cloudFunction

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type parameter struct { }

	func (p *parameter) Date(param string) time.Time {

		myTime, err := time.Parse("2006-01-02", os.Getenv(param))

		if err != nil {

			return time.Time{}
		}

		return myTime
	}

	func (p *parameter) Time(param string) time.Time {

		myTime, err := time.Parse("2006-01-02 15:04:05", os.Getenv(param))

		if err != nil {

			return time.Time{}
		}

		return myTime
	}

	func (p *parameter) Float(param string) float64 {

		number, err := strconv.ParseFloat(os.Getenv(param), 64)

		if err != nil {

			return 0
		}

		return number
	}

	func (p *parameter) Int(param string) int {

		myInt, err := strconv.Atoi(os.Getenv(param))

		if err != nil {

			myInt = 0
		}

		return myInt
	}

	func (p *parameter) Bool(param string) bool {

		switch strings.ToLower(strings.TrimSpace(os.Getenv(param))) {

			case "yes", "true", "on", "positive":

				return true
		}

		return false
	}

	func (p *parameter) String(param string) string {

		return os.Getenv(param)
	}

	func (p *parameter) Exists(param string) bool {

		if len(os.Getenv(param)) < 1 {

			return false
		}

		return true
	}


