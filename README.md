# RECORDMAKER

## Flattens the contents of multiple lines between a record seperator  into one line

```
/recordmaker_mac
Usage: ./recordmaker_mac -file <inputFile> -rs <recordSeparator> [-output <outputFile>] [-ls <lineSeparator>] [-rstop <recordStop>]

Options:
  -file string
    	Path to input file
  -ls string
    	Line separator. Defines the character between each newline of the record that has been flattened into a single line (optional, must be in quotes)
  -output string
    	Path to output file (optional)
  -rs string
    	Record separator. Define where a record starts (required, must be in quotes)
  -rstop string
    	Defines where the record stops. Needs to be in quotes (optional)
```
