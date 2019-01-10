package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Let's do some invitation figuring-outting!")

	GetPartners()
}

var getUrl = "https://candidate.hubteam.com/candidateTest/v3/problem/dataset?userKey=ccbae50eb46fe5cc58ff2d925903"

//This struct will represent every element in the GET response list.
type Partner struct {
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Email          string   `json:"email"`
	Country        string   `json:"country"`
	AvailableDates []string `json:"availableDates"`
}

//This struct is for the entire list response.
type Partners struct {
	Partners []Partner `json:"partners"`
}

func GetPartners() {
	fmt.Println("About to issue GET request to the hubteam server at URL: ", getUrl)

	getResp, err := http.Get(getUrl)

	if err != nil {
		//Woops. Something went wrong.
		println("Flagrant error! Couldn't GET that URL properly.")
		os.Exit(1)
	}
	defer getResp.Body.Close()

	if getResp.StatusCode == http.StatusOK {
		fmt.Println("We successfully got something!")
	}
}
