package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("please provide 1 argument, it should be the amount of seconds to wait")
		os.Exit(1)
	}
	amount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	sleep(time.Duration(amount) * time.Second)
	fmt.Println("done!")
}

func sleep(amount time.Duration) {
	<-time.After(amount)
}
