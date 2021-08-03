package olapdb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lbryio/lbrytv/apps/watchman/gen/reporter"
	"github.com/lbryio/lbrytv/apps/watchman/log"
	"github.com/pkg/errors"

	_ "github.com/ClickHouse/clickhouse-go"
)

var (
	conn     *sql.DB
	database string
)

func Connect(url string, dbName string) error {
	var err error
	if dbName == "" {
		dbName = database
	}
	database = dbName
	conn, err = sql.Open("clickhouse", url)
	if err != nil {
		return err
	}

	go ping()

	_, err = conn.Exec(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %v`, dbName))
	if err != nil {
		return err
	}
	_, err = conn.Exec(fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %v.playback
	(
		"URL" String,
		"Duration" UInt32,
		"Timestamp" Timestamp,
		"Position" UInt32,
		"RelPosition" UInt8,
		"RebufCount" UInt8,
		"RebufDuration" UInt32,
		"Protocol" FixedString(3),
		"Cache" String,
		"Player" FixedString(16),
		"UserID" UInt32,
		"Bandwidth" UInt32,
		"Device" FixedString(3),
		"Area" FixedString(2),
		"SubArea" FixedString(3),
		"IP" IPv6
	)
	ENGINE = MergeTree
	ORDER BY (Timestamp, UserID, URL)
	TTL Timestamp + INTERVAL 30 DAY`, dbName))
	if err != nil {
		return err
	}
	log.Log.Named("clickhouse").Infof("connected to clickhouse server %v (database=%v)", url, dbName)
	return nil
}

func Write(stmt *sql.Stmt, r *reporter.PlaybackReport, addr string, ts string) error {
	// t, err := time.Parse(time.RFC1123Z, r.T)
	var (
		t     time.Time
		err   error
		rate  uint32
		cache string
	)
	if ts != "" {
		t, err = time.Parse(time.RFC1123Z, ts)
		if err != nil {
			return err
		}
	} else {
		t = time.Now()
	}
	area, subarea := getArea(addr)
	// if area == "" {
	// 	area = "unknown"
	// }

	if r.Bandwidth != nil {
		rate = uint32(*r.Bandwidth)
	}
	if r.Cache != nil {
		cache = (*r.Cache)
	} else {
		cache = "miss"
	}

	_, err = stmt.Exec(
		r.URL,
		uint32(r.Duration),
		t,
		uint32(r.Position),
		uint8(r.RelPosition),
		uint8(r.RebufCount),
		uint32(r.RebufDuration),
		r.Protocol,
		cache,
		r.Player,
		uint32(r.UserID),
		rate,
		r.Device,
		area,
		subarea,
		addr,
	)
	if err != nil {
		return err
	}
	log.Log.Named("clickhouse").Infow(
		"playback record written",
		"user_id", r.UserID, "url", r.URL, "rebuf_count", r.RebufCount, "area", area, "ip", addr, "ts", t)
	return nil
}

func WriteOne(r *reporter.PlaybackReport, addr string, ts string) error {
	tx, err := conn.Begin()
	if err != nil {
		return errors.Wrap(err, "cannot begin")
	}
	stmt, err := prepareWrite(tx)
	if err != nil {
		return errors.Wrap(err, "cannot prepare")
	}
	defer stmt.Close()
	Write(stmt, r, addr, ts)
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "cannot commit")
	}
	return nil
}

func prepareWrite(tx *sql.Tx) (*sql.Stmt, error) {
	return tx.Prepare(fmt.Sprintf(`
	INSERT INTO %v.playback
		(URL, Duration, Timestamp, Position, RelPosition, RebufCount,
			RebufDuration, Protocol, Cache, Player, UserID, Bandwidth, Device, Area, SubArea, IP)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, database))
}

func ping() {
	ticker := time.NewTicker(60 * time.Second)
	for range ticker.C {
		if err := conn.Ping(); err != nil {
			log.Log.Named("clickhouse").Errorf("error pinging clickhouse: %v", err)
		}
	}
}
