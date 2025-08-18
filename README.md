# API Interview

Included in this repository are three binaries, one for windows (`server-windows.exe`), one for intel mac (`server-osx-intel`), and one for m-series mac (`server-osx`). This binary will start a basic HTTP JSON API on port 12000. This HTTP server has a single endpoint `/api/v1/punches` that you can find documented [here](#apiv1punches).

# Part 1

We would like you collect all data from this API for the date range of 2025-02-01 to 2025-02-28.

The expected number of records for this date range is 338.

# Part 2

We would like you count how many hours were worked on 2025-02-01.

The expected number of hours is 43.

# Part 3

We would like you count how many days in February 2025 have no hours worked.

The expected number is 1.

## Running server

It is important that the server is started from the `server` directory in order for the server to find the `interview.db` sqlite file.

1. `cd server`
2. `./bin/server-osx`
    - or `./bin/server-osx-intel`
    - or `./bin/server-windows.exe`

## Routes

### /api/v1/punches

#### Query Params

- `from`: required, date in format of `YYYY-MM-DDTHH:MM:SSZ`
- `to`: required, date in format of `YYYY-MM-DDTHH:MM:SSZ`
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
