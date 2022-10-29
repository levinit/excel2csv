package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"path"
	"strings"
)

// ---------
var appName = "excel2csv"
var version = "0.1"
var (
	v          bool // version
	usage      bool
	outputDir  string
	inputFiles []string
	createDir  bool // create output directory if not exists
	delimiter  = ","
	force      bool
)

// ---------
func init() {
	curDir, _ := os.Getwd()
	flag.StringVar(&outputDir, "o", curDir, "output directory")
	flag.StringVar(&delimiter, "d", ",", "delimiter")
	flag.BoolVar(&createDir, "c", false, "create output directory without asking")
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
		if !createDir {
			var confirm string
			println(outputDir, "not exists, do you want to create it? (y/n)")
			if _, err := fmt.Scanln(&confirm); err != nil {
				log.Fatal(err)
			}
			if confirm != "y" {
				log.Fatal("output directory not exists")
			}
		}
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	for _, f := range inputFiles {
		log.Println("processing:", path.Base(f))
		readXlsx(f)
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

	if len(sheets) == 1 {
		rows, _ := f.GetRows(sheets[0])
		csvFileName := strings.TrimSuffix(path.Base(filepath), path.Ext(filepath)) + ".csv"
		saveCsv(rows, path.Join(outputDir, csvFileName))
		return
	}

	for _, sheet := range sheets {
		rows, _ := f.GetRows(sheet)

		csvFileName := strings.TrimSuffix(path.Base(filepath), path.Ext(filepath)) + "_" + sheet + ".csv"
		saveCsv(rows, path.Join(outputDir, csvFileName))
	}
}

func saveCsv(strList [][]string, filepath string) {
	println("generating:", path.Base(filepath))
	//check file exists
	if _, err := os.Stat(filepath); err == nil && !force {
		println(filepath, "exists, do you want to overwrite it? (y/n)")
		var confirm string
		if _, err := fmt.Scanln(&confirm); err != nil {
			log.Fatal(err)
		}
		if confirm != "y" {
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
