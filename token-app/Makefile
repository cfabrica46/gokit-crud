test: 
	@echo "Ejecutando tests token-app..."
	@go test ./... --cover

cover: 
	@echo "Ejecutando tests token-app..."
	@go test ./... --coverprofile coverage.out
	@go tool cover -func coverage.out
