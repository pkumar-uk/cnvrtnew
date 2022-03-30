package main

import (
	"fmt"
	"os"

	"cnvrtnew/logger"
	"cnvrtnew/sport"

	"github.com/spf13/viper"
)

func main() {

	//check if the arguments are being correctly passed
	//sourcefile is required in csv format of a particular structure
	// format of the inptut file
	//Name, Age, Gender, Sport(list) name, height, weight
	//destinationfile is where the output will be stored as a YAML or JSON

	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("Attn*** Require two arguments: sourcefile destinationfile")
		logger.Fatal("Attn*** Require two arguments: sourcefile destinationfile")
		return
	}
	inputFile := arguments[1]
	outputFile := arguments[2]
	//get the environment valirable
	_ = viper.BindEnv("FORMAT")

	convertType := viper.Get("FORMAT")
	if !(convertType == "YAML" || convertType == "JSON") {
		logger.Fatal("Error no environment variable FORMAT set with YAML or JSON as value")
		return
	}

	//process the file
	slicelist, err := sport.ProcessFile(inputFile)
	if err != nil {
		logger.Fatal("Not able to process input file " + inputFile)
		return
	}

	//what do we want to convert to

	//convert to JSON
	if convertType == "JSON" {
		logger.Info("Will be converting Inputfile data to JSON")
		inputSlice := sport.DetailWriterJson{slicelist}
		sport.OutfileWriter(inputSlice, outputFile)
	}

	//convert to YAML
	if convertType == "YAML" {
		logger.Info("Will be converting Inputfile data to YAML")
		//create a placeholder and then convert
		inputSlice := sport.DetailWriterYaml{slicelist}
		sport.OutfileWriter(inputSlice, outputFile)
	}
}
