package main

import (
	"bufio"
	"fmt"
	//"io"
	//"io/ioutil"
	"os"
	"log"
	"strings"
	"regexp"
	"strconv"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
	//defer f.close()
}

func main() {
	//fmt.Printf("hello, world\n")

	file, err := os.Open("e:/FinCrime/epe27.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//var responseTimes []int
	var readLine =""
	var i int
	scanner := bufio.NewScanner(file)
	i=0
	var total int = 0
	var max int = 0
	var countSLA int = 0
	//responseTimes[0] = 0
	for scanner.Scan() {
		readLine = scanner.Text()
		//fmt.Println(i)
		if strings.Contains(readLine, "Orchestrator::Time taken for SAS processing"){
			var responseTime = regexp.MustCompile(`: ([0-9]+)`)
			fmt.Println(responseTime.FindStringSubmatch(readLine))
			var response string =   responseTime.FindStringSubmatch(readLine)[1]
			//var splitString = strings.Split(responseTime.FindStringSubmatch(readLine), " ")
			fmt.Println(response)
			var x int
			x, err = strconv.Atoi(response)
			total += x
			i++
			if x > 100 {
				countSLA++
			}
			if x > max {
				max = x
			}

		}
	}
	//var total int = 0


	var average float64 = float64(total) / float64(i)
	fmt.Printf("Average --> %.2f ms\n" , average)
	fmt.Printf("Max Value is %d\n", max)
	fmt.Printf("Count of transactions breaching SLA -> %d", countSLA)



	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
