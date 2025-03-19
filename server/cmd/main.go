package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type TimePunch struct {
	Id          int    `json:"id"`
	TimeClockID string `json:"time_clock_id"`
	PunchIn     string `json:"punch_in"`
	PunchOut    string `json:"punch_out"`
}

func main() {
	db, err := sql.Open("sqlite3", "./interview.db")

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/punches", func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		limit := r.URL.Query().Get("limit")

		if limit == "" {
			limit = "100"
		}

		offset := r.URL.Query().Get("offset")
		if offset == "" {
			offset = "0"
		}
		if from == "" || to == "" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"message": "from and to are required parameters",
			})
			return
		}
		qCount := `select
				count(id)
			 from time_punches
			 where
				 punch_in > ?
			 and
				 punch_out < ?`
		countRows, err := db.Query(qCount, from, to)
		if err != nil {
			log.Fatal(err)
		}
		defer countRows.Close()

		var count int

		for countRows.Next() {
			if err := countRows.Scan(&count); err != nil {
				log.Fatal(err)
			}
		}

		q := `select
				id,
				time_clock_id,
				punch_in,
				punch_out
			 from time_punches
			 where
				 punch_in > ?
			 and
				 punch_out < ?
			 order by punch_in asc
			 limit ?
			 offset ?`
		rows, err := db.Query(q, from, to, limit, offset)
		if err != nil {
			log.Fatal(err)
		}
		timePunches := []*TimePunch{}
		defer rows.Close()
		for rows.Next() {
			timePunch := &TimePunch{}
			err = rows.Scan(
				&timePunch.Id,
				&timePunch.TimeClockID,
				&timePunch.PunchIn,
				&timePunch.PunchOut,
			)
			if err != nil {
				log.Fatal(err)
			}
			timePunches = append(timePunches, timePunch)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"meta": map[string]any{
				"limit":        limit,
				"offset":       offset,
				"totalRecords": count,
			},
			"data": timePunches,
		})
	})

	s := &http.Server{
		Addr:           ":12000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Listening on port %v", 12000)

	log.Fatal(s.ListenAndServe())
}
