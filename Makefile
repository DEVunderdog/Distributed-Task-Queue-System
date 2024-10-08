DB_URL=postgresql://root:root@localhost:5432/distributed_queue_db?sslmode=disable

createdb:
	docker exec -it test-postgresql createdb --username=root --owner=root distributed_queue_db

dropdb:
	docker exec -it test-postgresql dropdb distributed_queue_db

migrate_create:
	migrate create -ext sql -dir $(DIRECTORY) -seq $(NAME)

migrateup:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

migrateup_version:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose up $(VERSION)

migratedown:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

migratedown_version:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" -verbose down $(VERSION)

migrate_force:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" force $(VERSION)

migrate_version:
	migrate -path $(MIGRATION_PATH) -database "$(DB_URL)" version


.PHONY: createdb dropdb migrate_create migrateup migrateup_version migratedown migratedown_version migrate_force migrate_version