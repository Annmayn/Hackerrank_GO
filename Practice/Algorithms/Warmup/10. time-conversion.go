package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
    /*
     * Write your code here.
     */
    var conv_hour int
    var str_hour string
    hour,_ := strconv.Atoi(s[:2])
    time := s[len(s)-2:]
    if time=="PM"{
        conv_hour = (hour+12)%24
        if hour==12{
            conv_hour=12
        }
        //fix formatting
        str_hour = fmt.Sprintf("%02d", conv_hour)
    }else if time=="AM"{
        str_hour = fmt.Sprintf("%02d", hour)
        if hour==12{
            str_hour = "00"
        }
    }
    
    result := str_hour + s[2:len(s)-2]
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer outputFile.Close()

    writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
