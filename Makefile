build-test:
	go build -ldflags="-s -w" -o gun && mv ./gun ./test_project/
