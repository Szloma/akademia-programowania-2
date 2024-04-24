package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//gui "github.com/grupawp/warships-lightgui/v2"
)

type BoardZ struct {
	board       []string
	desc        string
	target_nick string
	wpbot       bool
}

func getXAuthToken() string {
	posturl := "https://go-pjatk-server.fly.dev/api/game"

	GameBoard := map[string]interface{}{
		"coords": []string{
			"A1",
			"A3",
			"B9",
			"C7",
			"D1",
			"D2",
			"D3",
			"D4",
			"D7",
			"E7",
			"F1",
			"F2",
			"F3",
			"F5",
			"G5",
			"G8",
			"G9",
			"I4",
			"J4",
			"J8",
		},
		"desc":        "Pierwsza gra",
		"nick":        "Janusz",
		"target_nick": "",
		"wpbot":       true,
	}
	b, err := json.Marshal(GameBoard)
	if err != nil {
		fmt.Println(err)
		return "nil"
	}
	body := []byte(b)

	response, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	response.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(response)
	if err != nil {
		panic(err)
	}

	responseToken := res.Header.Get("X-Auth-Token")
	defer res.Body.Close()
	if responseToken == "" {
		return ""
	}
	return responseToken

}

type GameMap struct {
	board []string
}

func unmarshalToSlice(body []byte) []string {
	var data map[string]string

	err := json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}
	var values []string
	for _, v := range data {
		values = append(values, v)
	}

	return values
}

func getBoard(token string) []string {
	req, err := http.NewRequest("GET", "https://go-pjatk-server.fly.dev/api/game/board", nil)

	req.Header.Add("X-Auth-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	outputString := []byte(body)
	log.Println(outputString)

	var data map[string]interface{}

	err = json.Unmarshal(outputString, &data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	log.Println("---------")
	fmt.Printf("%s", data["board"])
	log.Println("---------")
	return unmarshalToSlice(outputString)
}

func getGameStatus(token string) {
	req, err := http.NewRequest("GET", "https://go-pjatk-server.fly.dev/api/game", nil)

	req.Header.Add("X-Auth-Token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
}

func main() {
	token := getXAuthToken()
	fmt.Print("X-Auth-Token: ", token)
	getGameStatus(token)

	gameBoard := getBoard(token)
	fmt.Println("gameboard")
	fmt.Println(gameBoard)
	//board := gui.New(gui.NewConfig())
	//board.Import(gameBoardRaw)
	//board.Display()

	time.Sleep(1 * time.Second)

	//sb = newFunction()
	//log.Printf(sb)
}

func newFunction() string {
	resp, err := http.Get("https://go-pjatk-server.fly.dev/api/game")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb
}
