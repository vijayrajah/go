package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type data struct {
	Name  string `yaml:"name"`
	Rg    string `yaml:"rg"`
	Probe []struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
	} `yaml:"probe"`
	Rules []struct {
		Name  string `yaml:"name"`
		Port  int    `yaml:"port"`
		Probe string `yaml:"probe"`
	} `yaml:"rules"`
}

func main() {

	var data1 data

	yamlFile, err := ioutil.ReadFile("inp.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &data1)

	fmt.Println(data1)

	for idx := range data1.Probe {
		fmt.Println(data1.Probe[idx].Name)
	}

	fmt.Println("# of rules is ", len(data1.Rules))
	for idx := range data1.Rules {
		fmt.Println(data1.Rules[idx].Name)
	}

}
