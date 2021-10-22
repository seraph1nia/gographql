https://github.com/99designs/gqlgen

## Stappen
- go mod init github.com/Seraph1nia/gographql
- go get github.com/99designs/gqlgen
- gooi schema's in schema.graphqls
- go run github.com/99designs/gqlgen init
- go mod tidy # dit verzorgt dependencies
- gqlgen generate #genereert opnieuw op basis van de schema's