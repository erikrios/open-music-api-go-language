migrate \
  -path "${PWD}" \
  -database "postgres://developer:erikrios@localhost:5432/openmusicgo?sslmode=disable" \
  down
