package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the divisibleSumPairs function below.
func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
    var cnt int32
    for i:=0; i<len(ar); i+=1{
        for j:=i+1; j<len(ar); j+=1{
            if (ar[i]+ar[j])%k == 0{
                cnt+=1
            }
        }
    }
    return cnt
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nk := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := divisibleSumPairs(n, k, ar)

    fmt.Fprintf(writer, "%d\n", result)

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

