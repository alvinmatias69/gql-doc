gql-doc:
	@go mod download
	@go build

.PHONY: install
install:
	@install gql-doc /usr/local/bin/gql-doc
	@mkdir -p /usr/local/lib/gql-doc/template
	@cp ./template/data/*.tmpl /usr/local/lib/gql-doc/template

.PHONY: test
test:
	@go test ./...

.PHONY: remove
remove:
	@rm /usr/local/bin/gql-doc
	@rm -rf /usr/local/lib/gql-doc
