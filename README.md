https://github.com/99designs/gqlgen

## Stappen
- go install github.com/99designs/gqlgen@latest
- go mod init bitbucket.org/Seraph1nia/gographq
- gooi schema's in schema.graphqls
- go run github.com/99designs/gqlgen init
- go mod tidy # dit verzorgt dependencies
- gqlgen generate #genereert opnieuw op basis van de schema's