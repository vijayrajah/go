// package main

// import (
// 	"fmt"

// 	"github.com/jszwec/csvutil"
// )

// type Status uint8

// const (
// 	Unknown = iota
// 	Success
// 	Failure
// )

// func (s Status) MarshalCSV() ([]byte, error) {
// 	fmt.Println("marshall")
// 	switch s {
// 	case Success:
// 		return []byte("success"), nil
// 	case Failure:
// 		return []byte("failure"), nil
// 	default:
// 		return []byte("unknown"), nil
// 	}
// }

// type Job struct {
// 	ID     int
// 	Status Status
// }

// func main() {
// 	jobs := []Job{
// 		{1, Success},
// 		{2, Failure},
// 	}

// 	b, err := csvutil.Marshal(jobs)
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	fmt.Println(string(b))

// 	// Output:
// 	// ID,Status
// 	// 1,success
// 	// 2,failure
// }

package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/jszwec/csvutil"
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

func (ad Address) MarshalCSV() (data []byte, err error) {

	fmt.Println("Inside MarshalCSV")

	buf := &bytes.Buffer{}
	err = binary.Write(buf, binary.BigEndian, ad)
	if err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return

}

// func (rec record) MarshalCSV() (data []byte, err error) {

// 	fmt.Println("Inside MarshalCSV")
// 	// err = nil
// 	// data = []string{rec.Name, strconv.Itoa(rec.Age), strconv.Itoa(rec.address.doorNo)}
// 	// return

// 	buf := &bytes.Buffer{}
// 	err = binary.Write(buf, binary.BigEndian, rec)
// 	if err != nil {
// 		panic(err)
// 	}
// 	data = buf.Bytes()
// 	return
// }

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

}
