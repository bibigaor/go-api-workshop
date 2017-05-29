package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"net/http"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"pass"`
}

type WorkerConf struct {
	Port string
}

var (
	userChan   chan User
	resultChan chan string
	//url chan string
)

func main() {
	// Initialize channels
	id := "1810"
	userChan = make(chan User)
	resultChan = make(chan string)
	idwork := url.QueryEscape(id)
	url := fmt.Sprintf("http://localhost.3001/workers/validate?idworker=%s",idwork)

	
	go mainApiWorker()



}

func mainApiWorker() {

	//se realiza el request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}

// Se valida que el request sea exitoso
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	var record User


	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	//se obtiene el usuario y contraseña
	 Nombre := record.Name
	 Password := record.Password


}
