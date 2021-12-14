package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	AppId                        string `yaml:"app_id"`
	AppUid                       string `yaml:"app_uid"`
	SecretKey                    string `yaml:"secret_key,omitempty"`
	AppBaseUrl                   string `yaml:"app_base_url"`
	MoyskladVendorApiEndpointUrl string `yaml:"moysklad_vendor_api_endpoint_url"`
	MoyskladJsonApiEndpointUrl   string `yaml:"moysklad_json_api_endpoint_url"`
}

func NewConfig(file string) *Config {
	c := new(Config)
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	if err = yaml.Unmarshal(data, c); err != nil {
		log.Fatalln(err)
		return nil
	}
	return c
}
