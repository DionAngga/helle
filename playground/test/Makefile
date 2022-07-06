
# buat postgresql
migrateup:
migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/go_sample?sslmode=disable" -verbose up
migratedown:
migrate -path db/migration -database "postgresql://postgres@admin:5432/go_sample?sslmode=disable" -verbose down
# root:@tcp(localhost:3306)/helle?charset=utf8mb4&parseTime=True&loc=Local

# buat mysql
migrateup:
migrate -path db/migration -database "mysql://root:@localhost:3306/go_sample?charset=utf8mb4&parseTime=True&loc=Local" -verbose up
migratedown:
migrate -path db/migration -database "mysql://root:@localhost:3306/go_sample?charset=utf8mb4&parseTime=True&loc=Local" -verbose down