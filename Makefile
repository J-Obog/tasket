TEST_DIR_NAME = tests

gen-mocks:
	mockery --all --output ./$(TEST_DIR_NAME)/mocks

unit-tests:
	go test -v ./$(TEST_DIR_NAME)/unit