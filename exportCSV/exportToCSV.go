package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/jszwec/csvutil"
)

type record struct {
	Name    string
	Age     int
	Address string
}

func main() {

	addresses := make([]record, 0)

	for i := 0; i < 10; i++ {
		addresses = append(addresses, record{Name: "Vijay" + strconv.Itoa(i), Age: i + 30, Address: strconv.Itoa(i) + "Some random street"})
	}
	fmt.Println(addresses)

	csvData, err := csvutil.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	err = writeCsv("out,csv", csvData)
	if err != nil {
		panic(err)
	}

}

func writeCsv(file string, records []byte) (err error) {

	err = ioutil.WriteFile(file, records, 0600)
	if err != nil {
		return
	}

	return
	// filep, err := os.Create(file)
	// if err != nil {
	// 	return
	// }

	// defer filep.Close()

	// csvWriter := csv.NewWriter(filep)
	// defer csvWriter.Flush()

	// err = csvWriter.Write(records)
	// if err != nil {
	// 	return
	// }

}
