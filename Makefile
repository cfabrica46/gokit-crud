test: 
	@echo "Ejecutando tests..."
	@make -C ./app
	@make -C ./database-app
	@make -C ./token-app

test-app: 
	@echo "Ejecutando test app..."
	@make -C ./app

test-database: 
	@echo "Ejecutando test app..."
	@make -C ./database-app

test-token: 
	@echo "Ejecutando test app..."
	@make -C ./token-app

