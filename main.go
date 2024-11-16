package main

import (
	"bufio"
	"fmt"
	"github.com/niqt/go-lesson/webserver"
	"os"
)

func main() {
	for {
		fmt.Println("*** Menu Options ***")
		fmt.Println("1. FanIn")
		fmt.Println("2. FanIn with select")
		fmt.Println("3. Close channel after fixed time")
		fmt.Println("4. Close channel if the messages are slow")
		fmt.Println("5. Producer Consumer with mutex")
		fmt.Println("6. Producer Consumer with RW mutex")
		fmt.Println("7. Close channel using the close function")
		fmt.Println("8. Close channel sending message in a channel")
		fmt.Println("9. Sieve")
		fmt.Println("10. CPU intensive task")
		fmt.Println("11. Server http")
		fmt.Println("12. Panic example without recover")
		fmt.Println("13. Panic example with recover")
		fmt.Println("14. Hang hover (the application will be blocked)")
		fmt.Println("q. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')

		switch choice {
		case "q\n":
			return
		case "1\n":
			FanInExample()
		case "2\n":
			FanInWithSelectExample()
		case "3\n":
			FinishAfterFixedTime()
		case "4\n":
			FinishIfSlow()
		case "5\n":
			ProducerConsumer()
		case "6\n":
			ScoreBoardRW()
		case "7\n":
			CloseCallingClose()
		case "8\n":
			CloseWithChannel()
		case "9\n":
			Sieve()
		case "10\n":
			CPU()
		case "11\n":
			webserver.ServerHttp()
		case "12\n":
			PanicExampleWithoutRecover()
		case "13\n":
			PanicExampleWithRecover()
		case "14\n":
			ChanNil()
		}
	}
}
