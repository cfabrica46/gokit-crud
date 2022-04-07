test:
	@echo "Ejecutando tests app..."
	go test ./... --cover

cover:
	@echo "Ejecutando tests app..."
	go test ./... --coverprofile coverage.out
	go tool cover -func coverage.out

.PHONY: all clean test cover
