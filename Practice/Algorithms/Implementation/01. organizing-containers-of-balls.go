package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

// Complete the organizingContainers function below.
func organizingContainers(container [][]int32) string {
    //box capacity
    row_sum := make([]int, len(container), len(container))
    for i,box:=range container{
        for _,val:=range box{
            row_sum[i]+=int(val)
        }
    }
    //total balls
    column_sum := make([]int, len(container), len(container))
    for _,box:=range container{
        for i,val:=range box{
            column_sum[i]+=int(val)
        }
    }
    sort.Ints(row_sum[:])
    sort.Ints(column_sum[:])
    for i,num:=range row_sum{
        if num!=column_sum[i]{
            return "Impossible"
        }
    }
    return "Possible"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        var container [][]int32
        for i := 0; i < int(n); i++ {
            containerRowTemp := strings.Split(readLine(reader), " ")

            var containerRow []int32
            for _, containerRowItem := range containerRowTemp {
                containerItemTemp, err := strconv.ParseInt(containerRowItem, 10, 64)
                checkError(err)
                containerItem := int32(containerItemTemp)
                containerRow = append(containerRow, containerItem)
            }

            if len(containerRow) != int(int(n)) {
                panic("Bad input")
            }

            container = append(container, containerRow)
        }

        result := organizingContainers(container)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
