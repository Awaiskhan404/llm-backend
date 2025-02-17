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

.PHONY: pkg
pkg:
	@if [ -z "$(name)" ]; then \
		echo "Error: Provide a module name with 'make pkg name=<module>'"; \
		exit 1; \
	fi; \
	MODULE_PATH="pkg/$(name)"; \
	MODULE_NAME="$(name)"; \
	MODULE_NAME_CAPITALIZED="$$(echo $$MODULE_NAME | awk '{print toupper(substr($$1,1,1)) tolower(substr($$1,2))}')"; \
	echo "Creating new module: $$MODULE_NAME..."; \
	mkdir -p $$MODULE_PATH; \
	for TEMPLATE in controller model repository routes service mock_repository; do \
		OUTPUT_FILE="$$MODULE_PATH/$$MODULE_NAME_$${TEMPLATE}.go"; \
		cat scripts/templates/$$TEMPLATE.template | \
		sed "s/{{name}}/$$MODULE_NAME/g; s/{{Name}}/$$MODULE_NAME_CAPITALIZED/g" > $$OUTPUT_FILE; \
	done; \
	OUTPUT_FILE="$$MODULE_PATH/$$MODULE_NAME.go"; \
	cat scripts/templates/module.template | \
	sed "s/{{name}}/$$MODULE_NAME/g; s/{{Name}}/$$MODULE_NAME_CAPITALIZED/g" > $$OUTPUT_FILE; \
	echo "Module $$MODULE_NAME created successfully!"


.PHONY: test
test:
	go test $(if $(VERBOSE),-v,) ./pkg/...
