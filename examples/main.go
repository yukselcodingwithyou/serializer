package main

import (
	"fmt"
	"github.com/yukselcodingwithyou/serializer"
)

func main() {
	sr := serializer.New(serializer.JSON)

	res, err := sr.Serialize(struct {
		Name string
	}{Name: "yuksel"})

	if err != nil {
		return
	}
	fmt.Println(res)
}
