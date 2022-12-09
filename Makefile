TEST_DIR_NAME = tests

generate-mocks:
	mockery --all --output ./$(TEST_DIR_NAME)/mocks

unit-tests:
	go test -v ./$(TEST_DIR_NAME)/unit