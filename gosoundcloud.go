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

func processAndUnmarshalResponses(resp *http.Response, err error, holder interface{}) error {
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	//TODO: check if StatusCode is 40x/50x if so set the body as the error and return
	if err = json.NewDecoder(resp.Body).Decode(holder); err != nil {
		return err
	}
	return nil
}

func processDeleteResponses(resp *http.Response, err error) error {
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
