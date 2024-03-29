package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
    var min, max int32
    var sum int64 
    sum = int64(arr[0])
    min = arr[0]
    max = arr[0]
    for i:=1; i<len(arr); i++{
        if arr[i]<min{
            min=arr[i]
        }
        if arr[i]>max{
            max=arr[i]
        }
        sum+=int64(arr[i])
    }
    fmt.Printf("%d %d",sum-int64(max), sum-int64(min))
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
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
