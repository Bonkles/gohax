package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	var respBodyString string
	if getResp.StatusCode == http.StatusOK {
		fmt.Println("We successfully got something!")
		getRespBody, bodyParseErr := ioutil.ReadAll(getResp.Body)

		respBodyString = string(getRespBody)

		if bodyParseErr != nil {
			fmt.Printf("Something went wrong while parsing the body of the request. HTTP response object raw format: %+v", getResp)
			os.Exit(2)
		}
		//fmt.Println("GET body is as follows: ", respBodyString)
	}

	//Allocate some memories for what we're gonna unmarshal.
	var partnerResponseList Partners
	bodyBytes := []byte(respBodyString)

	//Convert the body bytestream into our object instance so that we can operate on it.
	unmarshalErr := json.Unmarshal(bodyBytes, &partnerResponseList)

	if unmarshalErr != nil {
		fmt.Println("Error occurred while unmarshalling the JSON into an object. Bailing out.")
		os.Exit(3)
	}

	for index, partner := range partnerResponseList.Partners {
		fmt.Printf("Partner %d:\n %+v\n", index, partner)
	}
}
