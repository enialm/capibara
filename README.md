<img align="right" width="120px" src="https://messier.ch/assets/capibara.png">

[![blazingly fast](https://blazingly.fast/api/badge.svg?repo=enialm%2Fcapibara)](https://blazingly.fast)

# Capibara

Capibara is an event counting API. It lets you quickly measure how often something happens. Common use cases include monitoring, product analytics, and more.

Capibara is built with Go, Gin, Postgres, and restraint.

## Features

- Record events by name via HTTP POST
- Retrieve event counts, optionally filtered by time range
- Delete all records for a specific event
- Delete all records for all events
- API key authentication

## Setup

### Container

Run the Docker container with the following environment variables:

- `DSN`: Postgres connection string
- `API_KEY`: API key required in requests
- `GIN_MODE`: Use `release` in production

### Database

Create the `events` table:

```sql
CREATE TABLE events (
  id SERIAL PRIMARY KEY,
  event TEXT NOT NULL,
  ts BIGINT NOT NULL
);

CREATE INDEX idx_events_ts ON events(ts);
```

## Endpoints

### Authentication and Headers

Requests must include an `X-API-Key` header containing the configured API key. Endpoints that accept a JSON body (`/event`, `/delete`) should be called with a `Content-Type` header set to `application/json`.

### Record Event

`POST /event`

Records an event occurrence with the provided name. The API automatically adds a Unix timestamp (in seconds) corresponding to the ingestion time.

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

Returns recorded event counts by event name. The optional query parameters `start` and `end` can be used to filter results by Unix timestamps (in seconds), allowing custom aggregation. The bounds are inclusive, and either parameter may be omitted.

**Response:**  
```json
{
  "event_name_1": 128,
  "event_name_2": 64
}
```

### Delete Records by Name

`POST /delete`

Deletes all records associated with the specified event name. Other events are unaffected.

**Body:**  
```json
{
  "event": "event_name"
}
```

**Response:**  
```json
{
  "status": "42 matching record(s) deleted"
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

Returns a simple response for health checks. This endpoint does not require authentication.

**Response:**  
```json
{
  "message": "pong"
}
```

## Contributing

This software is considered complete. If you require changes, please fork the repository.

## License

MIT (see [LICENSE](LICENSE))
