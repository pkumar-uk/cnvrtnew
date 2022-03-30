package sport

import (
	"bufio"
	"cnvrtnew/logger"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

// The structure of the each record in the inputfile
type PersonRecord struct {
	Name   string   `json:name`
	Age    int      `json:age`
	Gender string   `json:gender`
	Sport  []string `json:sport`
	Height float64  `json:height`
	Weight int      `json:weight`
}

//Interface definition
type DetailWriter interface {
	WriteToFile(file string) error
}

//Converter to JSON
type DetailWriterJson struct {
	List []PersonRecord
}

func (a DetailWriterJson) WriteToFile(file string) error {
	dat, err := json.MarshalIndent(a.List, "  ", "  ")
	if err != nil {
		logger.Error("Could not convert to JSON ")
		return err
	}
	err = os.WriteFile(file, dat, 0644)
	if err != nil {
		logger.Error("Problem writing to output file " + file)
		return err
	}
	return nil
}

//Converter to YAML
type DetailWriterYaml struct {
	List []PersonRecord
}

func (a DetailWriterYaml) WriteToFile(file string) error {
	dat, err := yaml.Marshal(a.List)
	if err != nil {
		logger.Error("Could not convert to YAML ")
		return err
	}
	err = os.WriteFile(file, dat, 0644)
	if err != nil {
		logger.Error("Problem writing to output file " + file)
		return err
	}
	return nil
}

func OutfileWriter(z DetailWriter, outfile string) {
	err := z.WriteToFile(outfile)
	if err != nil {
		logger.Error("Problem creating file " + outfile)
		return
	}
}

//process file and return a slice of PersonRecord

func ProcessFile(inputFile string) ([]PersonRecord, error) {

	logger.Info("Entered Processing file ")
	var personList []PersonRecord

	//try opening a file
	file, err := os.Open(inputFile)
	if err != nil {
		logger.Error("Problem reading input file " + inputFile)
		fmt.Println(err)
		return personList, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//read a line record
		rec := scanner.Text()
		//now we need to extract the sports field as it is a list as we cannot easily split the string by ','
		var person PersonRecord
		var recsplit []string
		var sportssplit []string
		//find the opening square bracket
		if leftpos := strings.Index(rec, "["); leftpos > -1 {
			//extract the string left of open square bracket
			leftpart := fmt.Sprintf("%v", rec[:leftpos])
			//Now get the position of closing square bracket and extract the right part of the string
			rightpos := strings.Index(rec, "]")
			rightpart := fmt.Sprintf("%v", rec[rightpos+2:])
			//now get the sports list
			sportlist := fmt.Sprintf("%v", rec[leftpos+1:rightpos])
			//rebuild the string keeping the sport field null
			recsplit = strings.Split(leftpart+","+rightpart, ",")
			sportssplit = strings.Split(sportlist, ",")
		} else {
			logger.Info("Sport field/column has no value " + rec)
			recsplit = strings.Split(rec, ",")
		}
		//populate the struct
		person.Name = recsplit[0]

		person.Age, err = strconv.Atoi(recsplit[1])
		if err != nil {
			logger.Info("Problem converting age setting it to 0 " + recsplit[0] + recsplit[1])
			person.Age = 0
		}

		person.Gender = recsplit[2]

		person.Sport = sportssplit

		person.Height, err = strconv.ParseFloat(recsplit[4], 64)
		if err != nil {
			logger.Info("Problem converting height setting it to 0 " + recsplit[0] + recsplit[4])
			person.Height = 0.0
		}

		person.Weight, err = strconv.Atoi(recsplit[5])
		if err != nil {
			logger.Info("Problem converting weight setting it to 0 " + recsplit[0] + recsplit[5])
			person.Weight = 0
		}

		//append the record to the slice
		personList = append(personList, person)
	}
	logger.Info("Exiting Processing file ")
	return personList, nil
}
