tests: 
	@echo "Ejecutando tests..."
	@make test-app
	@make test-database
	@make test-token

test-app: 
	@echo "Ejecutando test app..."
	@make -C ./app test

test-database: 
	@echo "Ejecutando test database..."
	@make -C ./database-app test

test-token: 
	@echo "Ejecutando test token..."
	@make -C ./token-app test


cover: 
	@echo "Ejecutando tests..."
	@make cover-app
	@make cover-database
	@make cover-token

cover-app: 
	@echo "Ejecutando test app..."
	@make -C ./app cover

cover-database: 
	@echo "Ejecutando test database..."
	@make -C ./database-app cover

cover-token: 
	@echo "Ejecutando test token..."
	@make -C ./token-app cover
