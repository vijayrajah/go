package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Record struct {
	Name string
	Age  int
	Address
}

type Address struct {
	DoorNo int
	Street string
	City   string
}

func writeCSV(f string, r []Record) (err error) {
	w, err := os.Create(f)
	if err != nil {
		panic(err)
	}
	c := csv.NewWriter(w)
	defer c.Flush()

	for key, a := range r {
		fmt.Println(key, " & A is ", a)
	}

	for _, row := range r {
		var data []string
		data = append(data, row.Name)
		data = append(data, strconv.Itoa(row.Age))
		data = append(data, strconv.Itoa(row.Address.DoorNo))
		data = append(data, row.Address.Street)
		data = append(data, row.Address.City)

		//finally write the data
		if err := c.Write(data); err != nil {
			fmt.Println("Unable to write CSV data..")
			panic(err)
		}
	}

	return
}

func main() {

	addresses := make([]Record, 0)

	for i := 0; i < 10; i++ {

		record := Record{
			Name: "Vijay" + strconv.Itoa(i),
			Age:  i + 30,
			Address: Address{
				DoorNo: i,
				Street: "Dummy street",
				City:   "Dumy city",
			},
		}

		addresses = append(addresses, record)
	}
	fmt.Println(addresses)

	writeCSV("out.csv", addresses)

}
