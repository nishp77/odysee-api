package olapdb

import (
	"fmt"
	"testing"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/lbryio/lbrytv/apps/watchman/gen/reporter"
	"github.com/lbryio/lbrytv/apps/watchman/log"
	"github.com/stretchr/testify/suite"
)

type batchWriterSuite struct {
	BaseOlapdbSuite
}

func TestBatchWriterSuite(t *testing.T) {
	log.Configure(log.LevelInfo, log.EncodingConsole)
	suite.Run(t, new(batchWriterSuite))
}

func (s *batchWriterSuite) TestBatch() {
	days := 14
	number := 1000
	reports := []*reporter.PlaybackReport{}
	bw := NewBatchWriter(1*time.Second, 16)
	go bw.Start()

	for t := range timeSeries(number, time.Now().Add(time.Duration(-days)*24*time.Hour)) {
		r := PlaybackReportFactory.MustCreate().(*reporter.PlaybackReport)
		ts := t.Format(time.RFC1123Z)
		err := bw.Write(r, randomdata.StringSample(randomdata.IpV4Address(), randomdata.IpV6Address()), ts)
		s.Require().NoError(err)
		reports = append(reports, r)
	}

	time.Sleep(10 * time.Second)
	bw.Stop()

	var (
		url      string
		duration int32
	)
	rows, err := conn.Query(fmt.Sprintf("SELECT URL, Duration from %s.playback ORDER BY Timestamp DESC", database))
	s.Require().NoError(err)
	defer rows.Close()
	for i, r := range reports {
		n := rows.Next()
		s.Require().True(n, "only %v rows in db, expected %v", i, len(reports))
		err = rows.Scan(&url, &duration)
		s.Require().NoError(err)
		s.Equal(r.URL, url)
		s.Equal(r.Duration, duration)
	}
	s.False(rows.Next())
}
