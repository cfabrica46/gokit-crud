test:
	@echo "Ejecutando tests database-app..."
	go test ./... --cover

cover:
	@echo "Ejecutando tests database-app..."
	go test ./... --coverprofile coverage.out
	go tool cover -func coverage.out

.PHONY: all clean test cover
