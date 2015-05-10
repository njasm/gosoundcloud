package gosoundcloud

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Saver interface {
	Save(s *SoundcloudApi) error
}

type Updater interface {
	Update(s *SoundcloudApi) error
}

type Deleter interface {
	Delete(s *SoundcloudApi) error
}

type Resourcer interface {
	GetId() uint64
	GetKind() string
	IsNew() bool
}

//processAndUnmarshalResponses to process GET PUT POST request's responses. if the response have statusCode != 200
//then the response body will be unmarshalled into an error string and returned.
func processAndUnmarshalResponses(resp *http.Response, err error, holder interface{}) error {
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(bytes))
	}

	if err = json.NewDecoder(resp.Body).Decode(holder); err != nil {
		return err
	}

	return nil
}

//processDeleteResponses to process DELETE request's responses. if the response have statusCode != 200
//then the response body will be unmarshalled into an error string and returned.
func processDeleteResponses(resp *http.Response, err error) error {
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		bytes, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(bytes))
	}

	return nil
}
