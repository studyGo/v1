package main

import (
    "encoding/binary"
    "bytes"
    "net"
    "fmt"
    "os"
)

func main () {
    if len(os.Args) < 2 {
        usage()
        return
    }

    switch os.Args[1] {
        case "get_cinema_code":
            get_cinema_code()
        case "sign_data":
            sign_data(os.Args[2])
        default:
            usage()
    }
}

func usage () {
    fmt.Println("netsale get_cinema_code\n")
    fmt.Println("netsale sign_data  data\n")
}

func get_cinema_code () {
    check_file_exists()
    buffer := bytes.NewBuffer(make([]byte, 0, 1024))

    packetSize := uint32(12)
    packetType := uint32(1)
    function := uint32(9)

    binary.Write(buffer, binary.LittleEndian, packetSize)
    binary.Write(buffer, binary.LittleEndian, packetType)
    binary.Write(buffer, binary.LittleEndian, function)

    uc, _ := net.Dial("unix", "/tmp/.SOCKET_NETSALE")

    uc.Write(buffer.Next(12))

    b := make([]byte, 200)
    uc.Read(b)
    fmt.Println(string(b[12:20]))

}

func sign_data (d string) {
    check_file_exists()
    buffer := bytes.NewBuffer(make([]byte, 0, 1024))
    data := []byte(d)

    dataLen := 12 + len(data) + 4

    packetSize := uint32(dataLen)
    packetType := uint32(1)
    function := uint32(8)
    dataL := uint32(len(data))

    binary.Write(buffer, binary.LittleEndian, packetSize)
    binary.Write(buffer, binary.LittleEndian, packetType)
    binary.Write(buffer, binary.LittleEndian, function)
    binary.Write(buffer, binary.LittleEndian, data)
    binary.Write(buffer, binary.LittleEndian, dataL)

    uc, _ := net.Dial("unix", "/tmp/.SOCKET_NETSALE")

    uc.Write(buffer.Next(dataLen))

    b := make([]byte, 200)
    uc.Read(b)
    fmt.Print(string(b[12:140]))

}

func check_file_exists() {
    _, err := os.Stat("/tmp/.SOCKET_NETSALE")
    if err != nil {
        fmt.Println("file /tmp/.SOCKET_NETSALE is not exists")
        os.Exit(0)
    }
}
