install:
	go mod tidy && go mod vendor && go mod verify

mock:
	mockgen -source dependency.go -destination dependency_test.go -package ots_test
