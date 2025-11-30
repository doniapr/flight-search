# flight-search

# How to run
```bash
  go get
  go mod tidy
  go run main.go
```

# Example cUrl
```
curl --location 'localhost:8000/flights/search' \
--header 'airlines;' \
--header 'price_start: 0' \
--header 'price_end: 1250000' \
--header 'sort_by: arrival asc' \
--header 'Content-Type: application/json' \
--data '{
    "origin": "CGK",
    "destination": "DPS",
    "departureDate": "2025-12-15",
    "passengers": 1,
    "cabinClass": "economy"
}'
```