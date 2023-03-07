package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	General struct {
		Admin string `json:"admin,omitempty"`
	} `json:"general,omitempty"`
	Server struct {
		Host string `json:"host,omitempty"`
		Port string `json:"port,omitempty"`
		Url  string `json:"url,omitempty"`
	} `json:"server,omitempty"`
	Database struct {
		Host string `json:"host,omitempty"`
		Port string `json:"port,omitempty"`
		User string `json:"username,omitempty"`
		Pass string `json:"password,omitempty"`
		Db   string `json:"db,omitempty"`
	} `json:"database,omitempty"`
	Gitlab struct {
		Host  string `json:"host,omitempty"`
		Token string `json:"token,omitempty"`
	} `json:"gitlab,omitempty"`
	LDAP struct {
		Host  string `json:"host,omitempty"`
		Pass  string `json:"pass,omitempty"`
		Field struct {
			Id        string `json:"id,omitempty"`
			FirstName string `json:"first_name,omitempty"`
			LastName  string `json:"last_name,omitempty"`
			Mail      string `json:"mail,omitempty"`
			LdapUID   string `json:"ldap_uid,omitempty"`
		} `json:"field,omitempty"`
		Domain struct {
			Base  string `json:"base,omitempty"`
			Query string `json:"query,omitempty"`
		} `json:"domain,omitempty"`
		Filter struct {
			Employees string `json:"employees,omitempty"`
			Students  string `json:"students,omitempty"`
		} `json:"filter,omitempty"`
	} `json:"ldap,omitempty"`
}

func Configure() *Configuration {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Could not open config file!")
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	var config Configuration
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Config file not formated correctly!")
	}
	return &config
}
