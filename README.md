# A URL shortener example written in Go

## Generate OpenAPI server

```bash
go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
oapi-codegen --config config.yaml api.yaml > api/api.gen.go
```