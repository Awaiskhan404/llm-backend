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

.PHONY: rollback
rollback:
	@echo "Loading environment variables..."
	@export $(shell xargs < configs/.env)
	@echo "Rolling back migrations on database: $(DATABASE_NAME)..."

	# Execute all migration scripts in order
	@for file in $(shell ls sql/*.sql | sort); do \
		echo "Rolling back $$file..."; \
		PGPASSWORD=$(DATABASE_PASSWORD) psql -h $(DATABASE_HOST) -U $(DATABASE_USERNAME) -d $(DATABASE_NAME) -f $$file; \
	done

	@echo "Rollback complete."

.PHONY: test
test:
	go test $(if $(VERBOSE),-v,) ./pkg/...



.PHONY: pkg
pkg:
	@if [ -z "$(name)" ]; then \
		echo "❌ Error: Provide a module name with 'make pkg name=<module>'"; \
		exit 1; \
	fi; \
	MODULE_NAME="$(name)"; \
	MODULE_PATH="pkg/$(name)"; \
	MODULE_NAME_CAPITALIZED="$$(echo $(name) | awk '{print toupper(substr($$1,1,1)) tolower(substr($$1,2))}')"; \
	echo "🚀 Creating new module: $$MODULE_NAME..."; \
	mkdir -p "$$MODULE_PATH" || { echo "❌ Failed to create directory $$MODULE_PATH"; exit 1; }; \
	\
	for FILE in controller model repository routes service module; do \
		if [ "$$FILE" = "module" ]; then \
			OUTPUT_FILE="$$MODULE_PATH/$$MODULE_NAME.go"; \
		else \
			OUTPUT_FILE="$$MODULE_PATH/$$MODULE_NAME""_$$FILE.go"; \
		fi; \
		echo "Generating $$OUTPUT_FILE"; \
		cat scripts/templates/$$FILE.template | sed "s/{{name}}/$$MODULE_NAME/g; s/{{Name}}/$$MODULE_NAME_CAPITALIZED/g" > "$$OUTPUT_FILE"; \
	done; \
	\
	echo "✅ Module $$MODULE_NAME created successfully!"
	echo "📝 Update the module routes in 'internal/bootstrap' to include the new module routes."
