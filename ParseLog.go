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
	"path/filepath"
	"io/ioutil"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
	//defer f.close()
}

func main() {
	//fmt.Printf("hello, world\n")
	//Read the list of files from directory
	var dirPath string = "E:/FinCrime/TestResults/0208_PeakLoad"
	var filePath string
	var writeFileName string
	var total int = 0
	var max int = 0
	var countSLA int = 0
	var i int
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	writeFileName = filepath.Join(dirPath,"result.txt")
	fileWrite, err := os.Create(writeFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileWrite.Close()
	for _, file := range files{
		filePath = filepath.Join(dirPath,file.Name())
		fmt.Println("Reading file --> " + filePath)


		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		//var responseTimes []int
		var readLine =""

		scanner := bufio.NewScanner(file)
		i=0

		//responseTimes[0] = 0
		for scanner.Scan() {
			readLine = scanner.Text()
			//fmt.Println(i)
			if strings.Contains(readLine, "Orchestrator::Time taken for SAS processing"){
				var responseTime = regexp.MustCompile(`: ([0-9]+)`)
				//fmt.Println(responseTime.FindStringSubmatch(readLine))
				var response string =   responseTime.FindStringSubmatch(readLine)[1]
				//var splitString = strings.Split(responseTime.FindStringSubmatch(readLine), " ")
				//fmt.Println(response)
				var x int
				x, err = strconv.Atoi(response)
				total += x
				i++
				var writeString string = strconv.Itoa(i) + "," + strconv.Itoa(x)
				fileWrite.WriteString(writeString +"\n")

				//Calculate max value
				if x > 100 {
					countSLA++
				}
				if x > max {
					max = x
				}

			}
		}
		fileWrite.Sync()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	//var total int = 0

	fmt.Printf("Total --> %d", total)
	fmt.Printf("Count --> %d", i);
	var average float64 = float64(total) / float64(i)
	fmt.Printf("Average --> %.2f ms\n" , average)
	fmt.Printf("Max Value is %d\n", max)
	fmt.Printf("Count of transactions breaching SLA -> %d", countSLA)
}
