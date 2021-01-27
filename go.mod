module go-app-resto

// +heroku goVersion go1.15
go 1.15

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.9.0
	github.com/pkg/errors v0.8.1
	github.com/rs/cors v1.7.0
	gorm.io/driver/postgres v1.0.6
	gorm.io/gorm v1.20.11
)
