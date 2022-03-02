SHELL = /bin/sh

.PHONY: mock
mock: $(call print-target)
	mockery --all

.PHONY: migrateup
migrateup: $(call print-target)
	migrate -path migrations -database "postgres://postgres:berkay1707@localhost:5432/todo?sslmode=disable" -verbose up
	
.PHONY: migratedown
migratedown: $(call print-target)
	migrate -path migrations -database "postgres://postgres:berkay1707@localhost:5432/todo?sslmode=disable" -verbose down

.PHONY: nodemongorun
nodemongorun: $(call print-target)
	nodemon --exec go run main.go --signal SIGTERM;   
	
define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef