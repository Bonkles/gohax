package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's do some invitation figuring-outting!")

	GetPartners()
}

var getUrl = "http://https://candidate.hubteam.com/candidateTest/v3/problem/dataset?userKey=ccbae50eb46fe5cc58ff2d925903"

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

}
