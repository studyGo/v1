
// server use udp

package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    service := ":8080";
    udpAddr, err := net.ResolveUDPAddr("udp4", service)
    checkError(err)
    conn, err := net.ListenUDP("udp", udpAddr)
    checkError(err)

    for {
        handleClient(conn)
    }

}
func handleClient(conn *net.UDPConn) {
    fmt.Println("handle")
    var buf [512]byte
    _, addr, err := conn.ReadFromUDP(buf[0:])
    fmt.Println(string(buf[0:]))
    if err != nil {
        return
    }
    daytime := time.Now().String()
    fmt.Println("write")
    conn.WriteToUDP([]byte("this is a demo"), addr)
    conn.WriteToUDP([]byte(daytime), addr)
    fmt.Println("write end")
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "error", err.Error())
        os.Exit(1)
    }
}

