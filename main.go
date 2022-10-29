package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path"
)

// ---------
var appName = "excel2csv"
var version = "0.1"
var (
	v          bool // version
	usage      bool
	outputDir  string
	inputFiles []string
	inputDir   string
	delimiter  = ","
	force      bool
)

// ---------
func main() {
	for _, f := range inputFiles {
		log.Println("processing", f)
		readXlsx(f)
	}
}

func init() {
	curDir, _ := os.Getwd()
	flag.StringVar(&outputDir, "o", curDir, "output directory")
	flag.StringVar(&delimiter, "d", ",", "delimiter")
	flag.StringVar(&inputDir, "i", "", "input directory")
	flag.BoolVar(&force, "f", false, "force to overwrite output file")
	flag.BoolVar(&v, "v", false, "version")
	flag.BoolVar(&usage, "h", false, "show help")
	flag.Usage = func() {
		println(appName, version, "\n")
		println("Usage:", path.Base(os.Args[0]), "[options] input.xlsx [input2.xlsx ...]")
		println("Options:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if usage {
		flag.Usage()
		os.Exit(0)
	} else if v {
		println("version:", version)
		os.Exit(0)
	}
	inputFiles = flag.Args()
	if len(inputFiles) == 0 {
		log.Fatal("no input file")
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		log.Fatal("output directory not exists")
	}
}

func readXlsx(filepath string) {
	f, err := excelize.OpenFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheets := f.GetSheetList()

	for _, sheet := range sheets {
		rows, err := f.GetRows(sheet)
		if err != nil {
			fmt.Println(err)
			return
		}
		saveCsv(rows, sheet+".csv")
	}
}

func saveCsv(strList [][]string, filepath string) {
	//check file exists
	if _, err := os.Stat(filepath); err == nil && !force {
		println(filepath, "exists, do you want to overwrite it? (y/n)")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}
		if input != "y" {
			println("skip overwrite:", filepath)
			return
		}
		println("overwrite:", filepath)
	}

	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	w := csv.NewWriter(f)
	w.Comma = rune(delimiter[0])
	err = w.WriteAll(strList)
	if err != nil {
		return
	}
}
