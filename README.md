<img align="right" width="120px" src="https://messier.ch/assets/capibara.png">

[![blazingly fast](https://blazingly.fast/api/badge.svg?repo=enialm%2Fcapibara)](https://blazingly.fast)

# Capibara

Capibara is an event counting API that lets you quickly assess how often something happens. It can be used for monitoring, product analytics, and more.

Capibara is built with Go, Gin, Postgres, and restraint.

## Features

- Record events by name via HTTP POST
- Retrieve event counts, optionally filtered by time range
- Delete all records for a specific event name
- Delete all records for all events
- API key authentication
- Fast and small

## Setup

### Container

Publish the Docker container and set the following environment variables:

- `DSN`: Postgres connection string
- `API_KEY`: API key required in requests
- `GIN_MODE`: Use `release` if using in production

### Database

Create the events table:

```sql
CREATE TABLE events (
  id SERIAL PRIMARY KEY,
  event TEXT NOT NULL,
  ts BIGINT NOT NULL
);

CREATE INDEX idx_events_ts ON events(ts);
```

## Endpoints

### Authentication

The header `X-API-Key` is used for authentication.

### Record Event

`POST /event`

Records an event occurrence with the provided name. The API automatically adds a Unix timestamp set to the time of ingestion.

**Body:**  
```json
{
  "event": "event_name"
}
```

**Response:**  
```json
{
  "status": "ok"
}
```

### Get Event Stats

`GET /stats?start=t1&end=t2`

Returns recorded event counts by event name. The optional query parameters `start` and `end` can be used to filter results by Unix timestamps.

**Response:**  
```json
{
  "event_name_1": 128,
  "event_name_2": 64
}
```

### Delete Records by Name

`POST /delete`

Deletes all records for the specified event name.

**Body:**  
```json
{
  "event": "event_name"
}
```

**Response:**  
```json
{
  "status": "42 matching records deleted"
}
```

### Delete All Records

`POST /truncate`

Deletes all records for all events.

**Response:**  
```json
{
  "status": "all records deleted"
}
```

### Health Check

`GET /ping`

Returns a simple response for health checks.

**Response:**  
```json
{
  "message": "pong"
}
```

## Contributing

This software is considered complete and finished. If you require changes, please fork the repository.

## License

MIT (see [LICENSE](LICENSE))
