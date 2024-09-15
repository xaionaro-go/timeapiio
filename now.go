package timeapiio

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type TimeResponseRaw struct {
	Year        uint   `json:"year"`
	Month       uint   `json:"month"`
	Day         uint   `json:"day"`
	Hour        uint   `json:"minute"`
	Second      uint   `json:"second"`
	Millisecond uint   `json:"milliSeconds"`
	DateTime    string `json:"dateTime"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	TimeZone    string `json:"timeZone"`
	DayOfWeek   string `json:"dayOfWeek"`
	DstActive   bool   `json:"dstActive"`
}

const TimeLayout = "2006-01-02T15:04:05.999999MST"

func (r *TimeResponseRaw) Parse() (time.Time, error) {
	if r == nil {
		return time.Time{}, nil
	}

	t, err := time.Parse(TimeLayout, r.DateTime+r.TimeZone)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse field `dateTime`: %w", err)
	}

	return t, nil
}

func (c TimeAPIIO) Now() (time.Time, error) {
	url := ptr(*c.BaseURL)
	url.Path += "/time/current/zone"
	url.RawQuery = "timeZone=UTC"
	resp, err := c.HTTPClient.Get(url.String())
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to query URL '%s': %w", url.String(), err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to read the body of the response from '%s': %w", url.String(), err)
	}

	var msg TimeResponseRaw
	if err := json.Unmarshal(body, &msg); err != nil {
		return time.Time{}, fmt.Errorf("unable to un-JSON-ize the response '%s' from '%s': %w", body, url.String(), err)
	}

	result, err := msg.Parse()
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to parse the message %#+v: %w", msg, err)
	}

	return result, nil
}

func Now() (time.Time, error) {
	return DefaultTimeAPIIO.Now()
}
