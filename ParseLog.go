package main

import (
	"bufio"
	"fmt"
	//"io"
	//"io/ioutil"
	"os"
	"log"
	"strings"
	//	"regexp"
	"strconv"
	"path/filepath"
	"io/ioutil"
	//"gonum/stats"


	"regexp"
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
	//var dirPath string = "E:\FinCrime\prodLogs"
	var dirPath string = "E:/FinCrime/prodLogs"
	var filePath string
	var writeFileName string
	//var total int = 0
	//var max int = 0
	//var countSLA int = 0
	var i int
	//var roListener float64        // Elapsed time prior to transaction processing, includes MQGet
	var dataEnricher int = 0        // Elapsed time of ODE internal overhead prior to MEH database fetch


	//var countRoListener int = 0
	var countDataEnricher int = 0

	//var totalCount float64 = 0.0
	var max int = 0;

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	writeFileName = filepath.Join(dirPath, "RoListener.txt")
	fileWrite, err := os.Create(writeFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileWrite.Close()
	for _, file := range files {
		filePath = filepath.Join(dirPath, file.Name())
		fmt.Println("Reading file --> " + filePath)

		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		//var responseTimes []int
		var readLine = ""

		scanner := bufio.NewScanner(file)
		i = 0

		//responseTimes[0] = 0
		for scanner.Scan() {
			readLine = scanner.Text()
			//fmt.Println(i)
			if strings.Contains(readLine, "Time taken by ROlistener to read"){ // Time Taken by data enricher
				var responseTime = regexp.MustCompile(`:===([0-9]+)`)
				//fmt.Println(responseTime.FindStringSubmatch(readLine))
				var response string =   responseTime.FindStringSubmatch(readLine)[1]
				//var splitString = strings.Split(responseTime.FindStringSubmatch(readLine), " ")
				//fmt.Println(response)
				var x int
				x, err = strconv.Atoi(response)
				//total += x
				i++
				var writeString string = strconv.Itoa(i) + "," + strconv.Itoa(x)
				fileWrite.WriteString(writeString +"\n")
				dataEnricher+= x

				//Calculate max value

				if x > max {
					max = x
				}

			}

			countDataEnricher = i

		}

		fileWrite.Sync()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	//var total int = 0

	//fmt.Printf("Total --> %d", total)
	//fmt.Printf("Count --> %d", i);
	var average float64 = float64(dataEnricher) / float64(countDataEnricher)
	fmt.Printf("Average --> %.2f ms\n" , average)
	fmt.Printf("Max Value is %d\n", max)
	//fmt.Println("Elapsed time prior to transaction processing, includes MQGet -> ", count66)

}


