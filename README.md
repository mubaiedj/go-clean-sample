# Go Clean Sample

Template to build microservices based in goland using clean architecture

## Config file example

```
web:
 port: 8080
feature:
 flags:
  login: true
  save: true
datasource:
  type: cockroachdb
  host: localhost
  port: 26257
  database:
  user:
  password:
jwt:
 key: <add-your-key>
 username: <add-your-username>
 password: <add-your-password>
 duration: <add-you-expiration-time>
```
## Run with go

```
go mod tidy
```

```
go run app/main.go
```

## Run with docker

```
docker build -t sample .
```

```
docker run -p 8080:8080 sample
```