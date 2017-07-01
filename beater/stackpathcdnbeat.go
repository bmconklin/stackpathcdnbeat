package beater

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/bmconklin/stackpathcdnbeat/config"
	"github.com/jmervine/go-maxcdn"
)

type Stackpathcdnbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
	api    *maxcdn.MaxCDN
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	logp.Info("Initializing beater.")
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}
	logp.Info("Config read successfully.")
	bt := &Stackpathcdnbeat{
		done:   make(chan struct{}),
		config: config,
	}

	// verify credentials and permission with a simple call.
	maxcdn.APIHost = config.Endpoint
	bt.api = maxcdn.NewMaxCDN(config.Credentials.Alias, config.Credentials.Key, config.Credentials.Secret)

	var data maxcdn.Generic
	resp, err := bt.api.Get(&data, "/zones.json", nil)
	if err != nil {
		return nil, fmt.Errorf("Error verfiying API credentials: %v", err)
	}
	logp.Info("API call made successfully.")
	var zonesResp struct {
		Zones []struct {
			Id string
		}
	}
	if err := json.Unmarshal(resp.Data, &zonesResp); err != nil {
		return nil, fmt.Errorf("Error reading response from API: %v", err)
	}
	for i, id := range strings.Split(config.Credentials.Sites, ",") {
		if i == 0 && id == "" {
			// empty field, nothing to do
			break
		}
		var foundZone bool
		for _, z := range zonesResp.Zones {
			if z.Id == id {
				foundZone = true
				break
			}
		}
		if !foundZone {
			return nil, fmt.Errorf("Could not find site %v, are you sure it's correct and you have permission?", id)
		}
	}

	logp.Info("Finished initializing beater.")
	return bt, nil
}

func (bt *Stackpathcdnbeat) Run(b *beat.Beat) error {
	logp.Info("stackpathcdnbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	qs := url.Values{}
	qs.Set("start", bt.config.Start)
	qs.Set("limit", "1000")                                 // no reason not to get the most possible, maybe config this later
	qs.Set("sort", "oldest")                                // start from the beginning and work forward
	end := time.Now().UTC().Add(24 * 365 * 100 * time.Hour) // unreasonably high as a default.
	if bt.config.End != "" {
		endTime, err := time.Parse(time.RFC3339, bt.config.End)
		if err != nil {
			return fmt.Errorf("Error when parsing end time: %v", err)
		}
		end = endTime
		qs.Set("end", bt.config.End)
	}
	if bt.config.Credentials.Sites != "" {
		qs.Set("zones", bt.config.Credentials.Sites)
	}
	startKey := ""

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		var logResp maxcdn.Logs
		fmt.Println(bt.config.Path)
		resp, err := bt.api.Request("GET", bt.config.Path, qs)
		if err != nil {
			logp.Err("Failed to get a response form the API.", err)
			return fmt.Errorf("Could not get a response from the API. %v", err)
		}
		err = json.NewDecoder(resp.Body).Decode(&logResp)
		resp.Body.Close()
		if err != nil {
			logp.Err("Failed to get a response form the API.", err)
			return fmt.Errorf("Could not get a response from the API. %v", err)
		}

		// done? stop after this iteration.
		if len(logResp.Records) == 0 && end.After(time.Now().UTC()) {
			logp.Info("Reached end_time specified. Stopping.")
			bt.Stop()
		}

		for _, l := range logResp.Records {
			timestamp, err := time.Parse(time.RFC3339, l.Time)
			fmt.Println(l)
			if err != nil {
				logp.Err("Unable to parse timestamp.", err)
				return fmt.Errorf("Unable to parse timestamp. %v", err)
			}
			event := common.MapStr{
				"@timestamp": common.Time(timestamp),
				"type":       b.Name,
				"data":       l,
			}
			bt.client.PublishEvent(event)
		}
		if logResp.NextPageKey != "" {
			startKey = logResp.NextPageKey
		}
		qs.Set("start_key", startKey)
	}
}

func (bt *Stackpathcdnbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
