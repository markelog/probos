package middleware

import (
	"os"

	"github.com/iris-contrib/middleware/jwt"
)

// Middlewares is the list of all supported middlewares
type Middlewares struct {
	JWT *jwt.Middleware
}

// Up sets up middlewares
func Up() *Middlewares {
	middlewares := &Middlewares{}
	secret := os.Getenv("JWT_SECRET")

	middlewares.JWT = jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return middlewares
}
