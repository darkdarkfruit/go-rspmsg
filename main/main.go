package main

import (
	"github.com/darkdarkfruit/go-rspmsg/rspmsg"
	"fmt"
)

func main() {
	rsp := rspmsg.NewS()
	fmt.Printf("%#s\n", rsp)
	data := make(rspmsg.M)
	rsp.Data = data
	fmt.Printf("%s\n", rsp)
	rsp.Data.(rspmsg.M)["hello"] = "world"
	fmt.Printf("%s\n", rsp)


}