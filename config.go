package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Product struct {
	Name    string `json:"name"`
	Source  string `json:"url"`
	Website string `json:"website,omitempty"`
	Image   string `json:"image"`
	Send    bool   `json:"send"`
	Description string `json:description`
}

type Partner struct {
	Name    string `json:"name"`
	Image   string `json:"image"`
	Description string `json:description`
	Discord string `json:"discord,omitempty"`
	Website string `json:"website,omitempty"`
	Send    bool   `json:"send"`
}

type RulesChannel struct {
	Id string `json:"id"`
}

type PartnersChannel struct {
	Id       string    `json:"id"`
	Partners []Partner `json:"partners,omitempty"`
}

type ProductsChannel struct {
	Id       string    `json:"id"`
	Products []Product `json:"products,omitempty"`
}

type Rules struct {
	Id 	string `json:id`
	Send bool `json:send`
	Rules string `json:rules`
}

type Config struct {
	Name		   string		   `json:name`
	Token          string          `json:"token"`
	Guild          string          `json:"guild"`
	MemberRole     string          `json:"memberRole"`
	WelcomeChannel string          `json:"welcomeChannel"`
	RulesChannel   string          `json:"rulesChannel"`
	ServerURL      string          `json:"serverURL"`
	Products       ProductsChannel `json:"products,omitempty"`
	Partners       PartnersChannel `json:"partners,omitempty"`
	Rules Rules `json:rules`
}

func LoadConfig() (Config, error) {
	var cfg Config
	file, err := os.Open("config.json")
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
