package main

import "log"

func main() {
	str := "-aaa-"
	withoutFirst := str[1:]
	withoutLast := str[:len(str)-1]

	log.Print(withoutFirst)
	log.Print(withoutLast)
}
