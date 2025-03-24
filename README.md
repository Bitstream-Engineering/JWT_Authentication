A JWT Authenticator in Golang

Packages used:
1. GORM
2. Gin
3. bcrypt
4. jwt-go
5. godotev


N/B - two unresolved bugs 
1. token contains an invalid number of segments
2. A bug where in using postman to test POST - http://127.0.0.1:8080/login, returns the error -> "error": "Invalid email or password"

