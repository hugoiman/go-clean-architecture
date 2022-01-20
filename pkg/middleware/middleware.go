package middleware

import "net/http"

type Middleware interface {
	Auth() http.HandlerFunc
}
