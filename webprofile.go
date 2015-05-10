package gosoundcloud

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type WebProfile struct {
	Id         uint64
	Kind       string
	Title      string
	Url        string
	Created_at string
	Service    string
	Username   string
}

func NewWebProfile() *WebProfile {
	return &WebProfile{
		Kind: "web-profile",
	}
}

func (wp *WebProfile) Delete(s *SoundcloudApi) error {
	resp, err := s.Delete(wp.Url)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode == 200 {
		return nil
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	return errors.New(string(bytes))
}

func (wp WebProfile) MarshalJSON() ([]byte, error) {
	j := map[string]map[string]interface{}{
		"web-profile": {
			"title":   wp.Title,
			"url":     wp.Url,
			"service": wp.Service,
		},
	}
	return json.Marshal(j)
}
