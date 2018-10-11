package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// TodoURL is the URL to connect to the server
var TodoURL string

// clientCreateTodo calls server to create a todo
func clientCreateTodo(tm *todoModel) (*todoModel, error) {
	gmj, err := json.Marshal(tm)
	if err != nil {
		Log.Printf("Error marshaling into JSON: %v", err)
		return nil, err
	}
	req, err := http.NewRequest("POST", TodoURL+"/todos", bytes.NewBuffer(gmj))
	if err != nil {
		Log.Printf("Error creating HTTP request: %v", err)
		return nil, err
	}
	Log.Printf("Client Greeting with URL %s", TodoURL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("Error sending HTTP request: %v", err)
		return nil, err
	}
	if res.StatusCode != 201 {
		Log.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return nil, errors.New("Response code not 201")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Log.Printf("Error reading response body: %v", err)
		return nil, err
	}
	var tmr todoModel
	err = json.Unmarshal(body, &tmr)
	if err != nil {
		Log.Printf("Error de-marshaling model from body: %v", err)
		return nil, err
	}
	Log.Printf("Created Greeting with id %v\n", tmr.ID)
	return &tmr, nil
}

// clientGetTodos returns all todos
func clientGetTodos() ([]todoModel, error) {
	req, err := http.NewRequest("GET", TodoURL+"/todos", nil)
	if err != nil {
		Log.Printf("Error creating HTTP request: %v", err)
		return nil, err
	}
	Log.Printf("Client GET with URL %s", TodoURL)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("Error sending HTTP request: %v", err)
		return nil, err
	}
	if res.StatusCode != 200 {
		Log.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return nil, errors.New("Response code not 200")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Log.Printf("Error reading response body: %v", err)
		return nil, err
	}
	var tms []todoModel
	err = json.Unmarshal(body, &tms)
	if err != nil {
		Log.Printf("Error de-marshaling model from body: %v", err)
		return nil, err
	}
	Log.Printf("Retrieved %d number of greetings\n", len(tms))
	return tms, nil
}

// clientGetTodo returns todo from ID
func clientGetTodo(ID uint) (*todoModel, error) {
	url := TodoURL + "/todos/" + strconv.Itoa(int(ID))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Log.Printf("Error creating HTTP request: %v", err)
		return nil, err
	}
	Log.Printf("Client GET with URL %s", url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("Error sending HTTP request: %v", err)
		return nil, err
	}
	if res.StatusCode != 200 {
		Log.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return nil, errors.New("Response code not 200")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Log.Printf("Error reading response body: %v", err)
		return nil, err
	}
	var tm todoModel
	err = json.Unmarshal(body, &tm)
	if err != nil {
		Log.Printf("Error de-marshaling model from body: %v", err)
		return nil, err
	}
	Log.Printf("Retrieved Greeting with id %v\n", tm.ID)
	return &tm, nil
}

// clientUpdateTodo calls server to update a todo
func clientUpdateTodo(tm *todoModel) error {
	tmj, err := json.Marshal(tm)
	if err != nil {
		Log.Printf("Error de-marshaling model from JSON: %v", err)
		return err
	}
	url := TodoURL + "/todos/" + strconv.Itoa(int(tm.ID))
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(tmj))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		Log.Printf("Error creating HTTP request: %v", err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("Error sending HTTP request: %v", err)
		return err
	}
	if res.StatusCode != 200 {
		Log.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return errors.New("Response code not 200")
	}
	Log.Printf("Updated Greeting with id %v\n", tm.ID)
	return nil
}

// clientDeleteTodo calls server to delete a todo
func clientDeleteTodo(ID uint) error {
	url := TodoURL + "/todos/" + strconv.Itoa(int(ID))
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		Log.Printf("Error creating HTTP request: %v", err)
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		Log.Printf("Error sending HTTP request: %v", err)
		return err
	}
	if res.StatusCode != 200 {
		Log.Printf("Unexpected HTTP response code %d", res.StatusCode)
		return errors.New("Response code not 200")
	}
	Log.Printf("Deleted Greeting with id %v\n", ID)
	return nil
}
