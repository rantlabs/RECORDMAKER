package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    // Define flags
    inputFile := flag.String("file", "", "Path to input file")
    outputFile := flag.String("output", "", "Path to output file (optional)")
    recordSeparator := flag.String("rs", "", "Record separator. Define where a record starts (required, must be in quotes)")
    lineSeparator := flag.String("ls", "", "Line separator. Defines the character between each newline of the record that has been flattened into a single line (optional, must be in quotes)")
    recordStop := flag.String("rstop", "", "Defines where the record stops. Needs to be in quotes (optional)")

    // Define custom help message
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -file <inputFile> -rs <recordSeparator> [-output <outputFile>] [-ls <lineSeparator>] [-rstop <recordStop>]\n", os.Args[0])
        fmt.Fprintln(flag.CommandLine.Output(), "\nOptions:")
        flag.PrintDefaults()
    }

    // Parse flags
    flag.Parse()

    // Validate required flags
    if *inputFile == "" || *recordSeparator == "" {
        flag.Usage()
        os.Exit(1)
    }

    // Read input file
    data, err := ioutil.ReadFile(*inputFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
        os.Exit(1)
    }

    // Split the content by the record separator
    content := string(data)
    records := strings.Split(content, *recordSeparator)

    // Create output
    var result strings.Builder
    for _, record := range records {
        trimmedRecord := strings.TrimSpace(record)
        if trimmedRecord != "" {
            // Define where the record stops
            if *recordStop != "" {
                stopPos := strings.Index(trimmedRecord, *recordStop)
                if stopPos != -1 {
                    trimmedRecord = trimmedRecord[:stopPos]
                }
            }
            // Replace newlines with the line separator if specified
            if *lineSeparator != "" {
                trimmedRecord = strings.ReplaceAll(trimmedRecord, "\n", *lineSeparator)
            } else {
                trimmedRecord = strings.ReplaceAll(trimmedRecord, "\n", " ")
            }
            // Write the record on one line, prefixed with the record separator and a space
            result.WriteString(fmt.Sprintf("%s %s\n", *recordSeparator, trimmedRecord))
        }
    }
    output := result.String()

    // Output to terminal or file
    if *outputFile == "" {
        fmt.Print(output)
    } else {
        err := ioutil.WriteFile(*outputFile, []byte(output), 0644)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
            os.Exit(1)
        }
    }
}
