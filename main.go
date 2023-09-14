package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

var file string
var toAddress string

func init() {
	flag.StringVar(&file, "filePath", "", "Give the full system path to the attachment.")
	flag.StringVar(&toAddress, "toAddress", "", "Give the e-mail address of the recipient.")
	flag.Parse()

	if file == "" || toAddress == "" {
		fmt.Println("Crucial parameters are missing.")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

type Config struct {
	SmtpServer   string `json:"smtpServer"`
	Port         int    `json:"port"`
	MailAddress  string `json:"MailAddress"`
	MailPassword string `json:"MailPassword"`
}

func LoadConfiguration(destinationFile string) Config {
	jsonFile, err := os.Open(destinationFile)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}

func SendMailAttechmemt(server string, port int, From string, password string, To string, filePath string) string {
	m := gomail.NewMessage()
	m.SetHeader("From", From)
	m.SetHeader("To", To)
	m.SetHeader("Subject", "Attachment!")
	m.SetBody("text/html", "Hello, Here is the attachment!")
	m.Attach(filePath)

	d := gomail.NewDialer(server, port, From, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	} else {
		return "Mail has been send."
	}
}

func main() {
	c := LoadConfiguration("conf.json")
	result := SendMailAttechmemt(c.SmtpServer, c.Port, c.MailAddress, c.MailPassword, toAddress, file)
	fmt.Println(result)
}
