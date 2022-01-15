package main

import (
	"errors"
	"fmt"
	"net/http"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForId(userId string) (string, bool) {
	name, ok := sds.userData[userId]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore {
		userData: map[string]string {
			"1": "Fred",
			"2": "Kevin",
			"3": "Harry",
		},
	}
}

type DataStore interface {
	UserNameForId(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string){
	lg(message)
}

type SimpleLogic struct {
	l Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userId string) (string, error) {
	sl.l.Log("in SayHello for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userId string) (string, error) {
	sl.l.Log("in SayGoodbya for " + userId)
	name, ok := sl.ds.UserNameForId(userId)
	if !ok {
		return "", errors.New("unknown user")
	}

	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic {
		l: l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	l Logger
	logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l: l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
