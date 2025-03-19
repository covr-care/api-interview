# API Interview

We would like you collect all data from this API for the date range of 2025-02-01 to 2025-02-28.

## Running server

1. `cd server`
2. `./bin/server-osx`
    - or `./bin/server-osx-intel`
    - or `./bin/server-windows.exe`

### /api/v1/punches

#### Query Params

- `from`: required, date in format of `YYYY-MM-DD`
- `to`: required, date in format of `YYYY-MM-DD`
- `limit`: optional, default 100
- `offset`: optional, default 0

#### Response

```json
{
  "data": [
    {
      "id": 998,
      "time_clock_id": "8f1e3ce7-7577-4df6-810e-c3b8aa92d4d6",
      "punch_in": "2025-01-01T22:15:40Z",
      "punch_out": "2025-01-02T00:15:40Z"
    }
  ],
  "meta": {
    "limit": "10",
    "offset": "10",
    "totalRecords": 1000
  }
}
```
