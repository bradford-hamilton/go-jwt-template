## Dependencies
- Be sure to have Go (1.19+)
- Be sure to have postgres running locally
___
## Database
- Set up development db with
  ```
  createdb go_jwt_template_dev
  ```
- Run migration:
  ```
  psql -U go_jwt_template_user -d go_jwt_template_dev -a -f internal/storage/migrations/schema.sql
  ```
___
## Usage
### Development
```
go run cmd/server/main.go
```
___
## Deployment

Build docker image after code changes:
```
docker build \
  --build-arg GO_JWT_TEMPLATE_SERVER_PORT={server_port} \
  --build-arg GO_JWT_TEMPLATE_DB_HOST={host} \
  --build-arg GO_JWT_TEMPLATE_DB_PORT={db_port} \
  --build-arg GO_JWT_TEMPLATE_DB_USER={db_username} \
  --build-arg GO_JWT_TEMPLATE_DB_PASSWORD={db_password} \
  --build-arg GO_JWT_TEMPLATE_DB_NAME={db_name} \
  --build-arg GO_JWT_TEMPLATE_SSL_MODE={ssl_mode} \
  -t bradfordhamilton/go_jwt_template:latest .
```

Push image:
```
docker push bradfordhamilton/go_jwt_template:latest
```
