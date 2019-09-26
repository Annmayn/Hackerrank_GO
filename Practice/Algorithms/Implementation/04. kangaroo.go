package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the kangaroo function below.
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
    if (x1>=x2 && v1>v2) || (x2>=x1 && v2>v1){
        //high speed and x1 at same or greater point than x2
        return "NO"
    }else if (x1>x2 && v1>=v2) || (x2>x1 && v2>=v1){
        //high or equal speed but ahead of another point already 
        return "NO"
    }else if x1==x2 && v1==v2{
        //same point, same speed
        return "YES"
    }else if (x2-x1)%(v1-v2)==0{
        //they meet at some point
        return "YES"
    }else{
        //they don't meet at any point in the future
        return "NO"
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    x1V1X2V2 := strings.Split(readLine(reader), " ")

    x1Temp, err := strconv.ParseInt(x1V1X2V2[0], 10, 64)
    checkError(err)
    x1 := int32(x1Temp)

    v1Temp, err := strconv.ParseInt(x1V1X2V2[1], 10, 64)
    checkError(err)
    v1 := int32(v1Temp)

    x2Temp, err := strconv.ParseInt(x1V1X2V2[2], 10, 64)
    checkError(err)
    x2 := int32(x2Temp)

    v2Temp, err := strconv.ParseInt(x1V1X2V2[3], 10, 64)
    checkError(err)
    v2 := int32(v2Temp)

    result := kangaroo(x1, v1, x2, v2)

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
