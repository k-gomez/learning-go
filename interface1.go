package main

import "fmt"

type LogicProvider struct{}

// this is the business logic
// it can be changed easily without changing the client
func (lp LogicProvider) Process(data string) string {
	return fmt.Sprintf("Processed data to: %s", data[1:])
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

// only the client is able to see the interface (Logic)
// the client gets the data and processes it using the LogicProvider
func (c Client) Program() {
	data := "hello world"
	// get data
	fmt.Println(c.L.Process(data))
}

func main() {
	// create a client with one LogicProvider
	// we can implement multiple business logics (LogicProvider)
	// this allows us to be independent on the logic we need
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
