TEST_DIR = ./tests
MOCKS_DIR = $(TEST_DIR)/mocks

gen-mocks:
	rm -rf $(MOCKS_DIR)
	mockery --all --output $(MOCKS_DIR)

unit-tests:
	go test -v $(TEST_DIR)/unit