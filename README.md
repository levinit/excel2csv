# About

Command line tool to convert excel to csv.

Download binary files in the Releases page.

# Usage

Run it in the terminal:

```shell
Usage: main [options] confirm.xlsx [input2.xlsx ...]
```


Options:

- `-c`		create output directory without asking

- `-f`		force to overwrite output file

- `-h`		show help

- `-d <Single_delimiter_character>`

  ​			delimiter (default ",")

- `-o <output_dir_path>`

  ​			output directory (default is current work directory)

- `-v`        show version

Example:

```shell
#treat 1.xlsx as confirm file, and output to current directory
./excel2csv 1.xlsx

#convert 1.xlsx and 2.xlsx to csv and output to the csv_dir directory
./excel2csv -o csv_dir 1.xlsx 2.xlsx

#create output directory without asking
./excel2csv -c output_dir 1.xlsx
```

If the xlsx file contains multiple sheets, multiple csv will be generated, and the name of the csv is xlsx file name + sheet name. For example, if the xlsx file name is `1.xlsx`, and it contains two sheets, `sheet1` and `sheet2`, then the output csv file name will be `1_sheet1.csv` and `1_sheet2.csv`. 