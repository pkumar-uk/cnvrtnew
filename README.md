# Introduction
cnvrtnew - GO Program to Convert csv file to JSON or YAML

The input file is expected to be in a certain format as given below

`Rohan,12,M,[cricket,football],5.9,50`  
`Rohit,11,M,[football],5.8,51`  
`Keerthi,13,F,[badminton,table tennis],5.5,45` 

Corresponding to the following fields

`Name, Age, Gender, Sport name, height, weight` 

*here **Sport name** is a list of sports, refer to in.txt file in the folder for data and format*

The output is an JSON/YAML Array with each element of the array corresponding to a record in the csv file

## Program structure

logger/logger.go - a wrapper around zap to make it easy to use from any package
sport/sport.go - the main logic is here. First the file is read and the records are converted to slice of recs(persons) and then converted to YAML or JSON as the case may be

## Prerequiste for running the program

You obviously need GO to be setup/installed on your machine

## Setting up the program
Go to the main folder of the project once it has been downloaded and unpacked on your machine

Issue the following command to download and setup packages on your system

`go mod tidy`

## Running the program

The program expects environment variable **FORMAT** to be set, as that is  used to determine if it is JSON conversion or YAML conversion. The acceptable values for FORMAT environment variable is **JSON** or **YAML**

Other than setting up of the environment variable, the program requires two arguments - 
- inputfile - csv file that needs to be converted, prefreably with full path
- outputfile - name of the outputfile with path where it needs to be stored

One can run the program by issiung command

`go run main.go inputfile outputfile`  

make sure to replace inputfile and outputfile with correct relevant filenames with path

### To build and run

To build run issue command

`go build`

this will creatre **cnvrtnew** an executable in the same folder, now run it by issuing the command

`./cnvrtnew inputfile outputfile`   

make sure to replace inputfile and outputfile with correct relevant filenames with path
## Testing
A small testing harness is created to test the processing of inputfile to convert to slice of Persons

The test depends on file in.txt go to folder cnvrtnew/sport

To run the test, issue the command

`go test  -v`

To run the benchmark test as well, issue the command

`go test  -v -bench=.`


