## Swagger instructions

### Requirements
- Check if https://github.com/swaggo/swag locally installed

#### If not 
 - go install github.com/swaggo/swag/cmd/swag@latest

 ### Update swagger
 - command swag init

 #### ERRORS
  - zsh: command not found -> PATH=$(go env GOPATH)/bin:$PATH

 ### URL
 - http://localhost:8081/api/swagger/index.html#/
