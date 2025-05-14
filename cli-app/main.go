package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {

	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		fmt.Println("crashing")
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}\s` + regexp.QuoteMeta(*level) + `\s([\w\s]+)`)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			fmt.Println(line)
		}
	}
}
