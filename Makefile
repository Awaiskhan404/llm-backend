include configs/.env
export $(shell sed 's/=.*//' configs/.env)

build:
	go build -o bin/api cmd/gin-server/main.go

run:
	go run cmd/gin-server/main.go



.PHONY: migrate
migrate:
	@echo "Loading environment variables..."
	@export $(shell xargs < configs/.env)
	@echo "Running migrations on database: $(DATABASE_NAME)..."

	# Execute all migration scripts in reverse order
	@for file in $(shell ls -r sql/*.sql | sort -r); do \
		echo "Executing $$file..."; \
		PGPASSWORD=$(DATABASE_PASSWORD) psql -h $(DATABASE_HOST) -U $(DATABASE_USERNAME) -d $(DATABASE_NAME) -f $$file; \
	done

	@echo "Migration complete."

.PHONY: test
test:
	go test $(if $(VERBOSE),-v,) ./pkg/...
