tests: test-app test-database test-token
	@echo "Running tests..."

test-app:
	make -C ./app test

test-database:
	make -C ./database-app test

test-token:
	make -C ./token-app test

# ---

cover: cover-app cover-database cover-token
	@echo "Running covers..."

cover-app:
	make -C ./app cover

cover-database:
	make -C ./database-app cover

cover-token:
	make -C ./token-app cover

# ---

lint: lint-app lint-database lint-token
	@echo "Running general golangci-lint..."

lint-app:
	make -i -C ./app lint

lint-database:
	make -i -C ./database-app lint

lint-token:
	make -i -C ./token-app lint


.PHONY: tests test-app test-database test-token cover cover-app cover-database cover-token lint lint-app lint-database lint-token all clean test
