# MUNI - Model University

## Start GraphQL server
```
go run cmd/main.go
```

## Add Ent
```
go run -mod=mod entgo.io/ent/cmd/ent new LabellingTask
```

## Generate files
```
go generate .
```

## DB

### Inspect Schema
```
atlas schema inspect \
  -u "postgresql://muni_app:$MUNI_DATABASE_PASSWORD@127.0.0.1/muni?search_path=public&sslmode=disable" \
  --format "{{ sql . }}"
```
