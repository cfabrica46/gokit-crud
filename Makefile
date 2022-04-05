tests: test-app test-database test-token
	@echo "Ejecutando tests..."

test-app:
	make -C ./app test

test-database:
	make -C ./database-app test

test-token:
	make -C ./token-app test

# ---

cover: cover-app cover-database cover-token
	@echo "Ejecutando covers..."

cover-app:
	make -C ./app cover

cover-database:
	make -C ./database-app cover

cover-token:
	make -C ./token-app cover

.PHONY: tests test-app test-database test-token cover cover-app cover-database cover-token all clean test
