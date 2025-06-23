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
    recordSeparator := flag.String("rs", "", "Record separator (required, must be in quotes)")

    // Define help message
    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -file <inputFile> -rs <recordSeparator> [-output <outputFile>]\n", os.Args[0])
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
            // Write the record on one line, prefixed with the record separator
            result.WriteString(fmt.Sprintf("%s%s\n", *recordSeparator, strings.ReplaceAll(trimmedRecord, "\n", " ")))
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
