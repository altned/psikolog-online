module github.com/example/psikolog-online

go 1.20

replace github.com/gin-gonic/gin => ./github.com/gin-gonic/gin

replace github.com/golang-jwt/jwt/v5 => ./github.com/golang-jwt/jwt/v5

require (
	github.com/gin-gonic/gin v0.0.0-00010101000000-000000000000
	github.com/golang-jwt/jwt/v5 v5.0.0-00010101000000-000000000000
)
