# Capibara

**Capibara** is an event counting API. It can be used in countless contexts, from monitoring to product analytics and more. It is built with Go, Gin, Postgres, and restraint.

## Features

- Record events by name via HTTP POST
- Retrieve event counts, optionally filtered by time range
- API key authentication
- Fast and small

## Setup

### Container

Publish the Docker container and set the following environment variables:

- `DSN`: Postgres connection string
- `API_KEY`: API key required in requests

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

### Add Event

`POST /event`

**Body:**  
```json
{
  "event": "event_name"
}
```

### Get Event Stats

`GET /stats?start=t1&end=t2`

The query parameters `start` and `end` can be used to filter results with unix timestamps.

**Response:**  
```json
{
  "event_name_1": 128,
  "event_name_2": 64
}
```

## Contributing

This software is considered complete and finished. Fork it if you want something else.

## License

MIT (see [LICENSE](LICENSE))
