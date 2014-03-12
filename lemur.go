package lemur

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	//"reflect"
	"strings"
	"time"
)

var (
	ErrTimeOut = errors.New("Request Time Out")
	ErrUnknown = errors.New("Unknow Error")
)

var DefaultPusher = NewPusher("")
var DefaultTimeOut = 5 * time.Second

func SetTimeOut(to time.Duration) {
	DefaultTimeOut = to
}

func SetApiKey(apikey string) {
	DefaultPusher.SetApiKey(key)
}

func NewPusher(token string) (r *Pusher) {
	return &Pusher{
		apiKey: token,
	}
}

type Pusher struct {
	apiKey string
	listId string
}

func (this *Pusher) SetApiKey(apiKey string) {
	this.apiKey = apiKey
}

func (this *Pusher) SetListId(listId string) {
	this.listId = listId
}

type ResultError struct {
	Type    string `json"type"`
	Message string `json"message"`
}

type ResultJson struct {
	Success bool `json:"success"`
	Error   *ResultError
}

type Result map[string]interface{}

type ChimpError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Name    string `json:"name"`
	Message string `json:"error"`
}

func (e ChimpError) Error() string {
	return fmt.Sprintf("%v: %v", e.Code, e.Message)
}

func (this *Pusher) Do(verb string, m interface{}, ech chan error) {
	data, err := json.Marshal(m)
	if err != nil {
		ech <- err
		return
	}

	resp, err := http.Post(getBaseURL(this.apiKey)+verb, "application/json", strings.NewReader(string(data)))
	if err != nil {
		ech <- err
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Lemur response:")
	fmt.Println(string(body))
	if err != nil {
		ech <- err
		return
	}
	rj := Result{}
	err = json.Unmarshal(body, &rj)
	if err != nil {
		ech <- err
		return
	}
	if rj["code"] != nil {
		e := &ChimpError{}
		err = json.Unmarshal(body, &e)
		if err != nil {
			ech <- err
			return
		}
		ech <- e
		return
	}
	ech <- nil
	return
}

func (this *Pusher) SendData(verb string, m map[string]interface{}) (err error) {
	if m["apikey"] == nil {
		m["apikey"] = this.apiKey
	}
	if m["id"] == nil {
		m["id"] = this.listId
	}

	ech := make(chan error)
	go this.Do(verb, m, ech)
	select {
	case err = <-ech:
		return
	case <-time.After(DefaultTimeOut):
		err = ErrTimeOut
		return
	}
	return
}
