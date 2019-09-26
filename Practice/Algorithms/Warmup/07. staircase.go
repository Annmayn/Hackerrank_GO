package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the staircase function below.
func drawSpace(n int32){
    var i int32
    for i=0; i<n; i++{
        fmt.Printf("%s"," ")
    }
}

func drawStair(n int32){
    var i int32
    for i=0; i<n; i++{
        fmt.Printf("%s","#")
    }
    fmt.Print("\n")
}

func staircase(n int32) {
    var i int32
    for i=1; i<=n; i++{
        drawSpace(n-i)
        drawStair(i)
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    staircase(n)
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
