package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
    min := scores[0]
    max := scores[0]
    counter := make([]int32,2,2)
    for i:=1;i<len(scores);i++{
        if scores[i]<min{
            counter[1]++
            min = scores[i]
        }
        if scores[i]>max{
            counter[0]++
            max = scores[i]
        }
    }
    return counter
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    scoresTemp := strings.Split(readLine(reader), " ")

    var scores []int32

    for i := 0; i < int(n); i++ {
        scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
        checkError(err)
        scoresItem := int32(scoresItemTemp)
        scores = append(scores, scoresItem)
    }

    result := breakingRecords(scores)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

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
