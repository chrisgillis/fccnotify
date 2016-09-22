package main

import (
	"fmt"
	"flag"
	"os"
	"net/http"
	"net/smtp"
	"encoding/json"
	"time"
)

type SearchResponse struct {
	Status string
	Errors SearchResponseErrorParent
	Licenses LicenseParent
}

type SearchResponseErrorParent struct {
	Err []SearchResponseError
}

type SearchResponseError struct {
	Code string
	Msg string
}

type LicenseParent struct {
	Page string
	RowPerPage string
	TotalRows string
	LastUpdate string
	License []License
}

type License struct {
	Name string `json:"licName"`
	Frn string
	Callsign string
	CategoryDesc string
	ServiceDesc string
	StatusDesc string
	ExpiredDate string
	LicenseId string 
	LicenseDetailUrl string `json:"licDetailURL"`
}

type EmailUser struct {
	User string
	Pass string
	Host string
	Port string
}

func main() {
	var usage = "Usage: fccnotify -frn <frn> [-m <minutes>]";

	var frn = flag.String("frn", "", "your ULS FRN identifier")
	var interval = flag.Int("m", 30, "number of minutes between checks, minimum 30, default 30")
	var gmailuser = flag.String("gmailaddr", "", "your gmail email address")
	var gmailpass = flag.String("gmailpass", "", "your gmail password")
	flag.Parse()

	if *frn == "" {
		fmt.Println("Invalid FRN specified.\n", usage)
		os.Exit(1)
	}

	if *interval < 30 {
		fmt.Println("Minutes must be >= 30\n", usage)
		os.Exit(1)
	}

	fmt.Println("Searching for FRN", *frn, "every", *interval, "minutes.")

	var url = "http://data.fcc.gov/api/license-view/basicSearch/getLicenses?format=json&searchValue=" + *frn
	
	for {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		searchResponse := new(SearchResponse)
		json.NewDecoder(resp.Body).Decode(searchResponse)

		if len(searchResponse.Errors.Err) > 0 {
			if searchResponse.Errors.Err[0].Code == "110" {
				fmt.Println("[", time.Now(), "]", "Nothing yet...")	
			} else {
				fmt.Println(searchResponse.Errors.Err[0].Msg)
			}
		} else {
			notificationTxt := "You got your callsign!\n\n"
			notificationTxt += "Name: " + searchResponse.Licenses.License[0].Name + "\n"
			notificationTxt += "FRN: " +  searchResponse.Licenses.License[0].Frn + "\n"
			notificationTxt += "Callsign: " +  searchResponse.Licenses.License[0].Callsign + "\n"
			notificationTxt += "Category: " +  searchResponse.Licenses.License[0].CategoryDesc + "\n"
			notificationTxt += "Service: " +  searchResponse.Licenses.License[0].ServiceDesc + "\n"
			notificationTxt += "Status: " +  searchResponse.Licenses.License[0].StatusDesc + "\n"
			notificationTxt += "Expires: " +  searchResponse.Licenses.License[0].ExpiredDate + "\n"
			notificationTxt += "LicenseID: " +  searchResponse.Licenses.License[0].LicenseId + "\n"
			notificationTxt += "LicenseURL: " +  searchResponse.Licenses.License[0].LicenseDetailUrl + "\n"
			
			fmt.Println("\n")
			fmt.Println(notificationTxt)

			if *gmailuser != "" && *gmailpass != "" {
				emailTxt := "From: " + *gmailuser + "\r\n"
				emailTxt += "To: " + *gmailuser + "\r\n"
				emailTxt += "Subject: You got your callsign!\r\n\r\n"
				emailTxt += notificationTxt

				emailUser := &EmailUser{*gmailuser, *gmailpass, "smtp.gmail.com", "587"}
				auth := smtp.PlainAuth("",emailUser.User, emailUser.Pass, emailUser.Host)
				err = smtp.SendMail(emailUser.Host+":"+emailUser.Port,
					auth,
					emailUser.User,
					[]string{*gmailuser},
					[]byte(emailTxt))
				if err != nil {
					fmt.Println(err)
					os.Exit(3)
				}
			}

			break
		}

		resp.Body.Close()

		time.Sleep(time.Minute * time.Duration(*interval))
	}
	
}