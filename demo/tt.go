package main

import (
	"fmt"

	"math/rand"

	"github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
)

type Person struct {
	Name string
	Age  int
}

var logger = logging.GetLogger()

func params(maps map[string]interface{}) {
	name := maps["name"].(string)
	age := maps["age"].(int)
	//logger.Debugf("hell world %s", name)
	logger.Debug("name is %s,age is %d", name, age)
}

func main() {
	fmt.Println("hello world")
	amp := make(map[string]interface{})
	amp["name"] = "james"
	amp["age"] = 23
	ttList := []string{"a", "b"}
	var randNum = len(ttList)

	x := rand.Intn(randNum)
	logger.Debug("x is %d", x)
	logger.Debug("v is %s", ttList[x])

	p := Person{
		Name: "james",
		Age:  23,
	}
	fmt.Printf("person dumps %+v", p)
}
