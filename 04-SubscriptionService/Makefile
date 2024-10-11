## env variables
# BINARY_NAME=myapp ## for mac or linux 
BINARY_NAME=myapp.exe ## for windows
DSN="postgres://postgres:password@localhost:5432/concurrency?sslmode=disable"
REDIS="127.0.0.1:6379"

## build: builds all binaries
build:
	@echo Building...
	@go build -o ${BINARY_NAME} ./cmd/web
	@echo Built!

## run: builds and runs the application
## for mac or linux 
# run: build
# 	@echo "Starting..."
# 	@env DSN=${DSN} REDIS=${REDIS} ./${BINARY_NAME} &
# 	@echo "Started!"
## for windows
run: build
	@echo Starting...
	powershell $$env:DSN='${DSN}'; $$env:REDIS='${REDIS}'; ./${BINARY_NAME}
	@echo Started!

## clean: runs go clean and deletes binaries
## for mac or linux 
# clean:
# 	@echo "Cleaning..."
# 	@go clean
# 	@rm ${BINARY_NAME}
# 	@echo "Cleaned!"
# for windows
clean:
	@echo Cleaning...
	@go clean
	-@DEL ${BINARY_NAME} 2>nul
	@echo Cleaned!

## start: an alias to run
start: run

## stop: stops the running application
## for mac or linux 
# stop:
# 	@echo "Stopping..."
# 	@-pkill -SIGTERM -f ./${BINARY_NAME}
# 	@echo "Stopped!"
## for windows
stop:
	@echo Stopping...
	-@taskkill /IM ${BINARY_NAME} /F 2>nul || echo Process not found or already stopped.
	@echo Stopped!

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	@echo Testing...
	go test -v ./...