package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "strings"
)

type Redis struct{
    conn *net.TCPConn
}

func New(ip string, port string) *Redis {
    conn := redis_connect(ip, port)
    return &Redis{
        conn : conn,
    }
}

func (r *Redis) Get(key string) (value string) {

    str := r.redis_cmd_format_static("get", key)
    r.write(str)
    str = r.read()
    return str

}

func (r *Redis) Set(key string, value string) bool {
    str := r.redis_cmd_format_static("set", key, value)
    r.write(str)
    str = r.read()
    if str == "OK" {
        return true
    } else {
        return false
    }
}

func (r *Redis) redis_cmd_format_static(args ... string) string {

    //*3\r\n$3\r\nSET\r\n$5\r\nmykey\r\n$7\r\nmyvalue\r\n
    str := "*"
    str += strconv.Itoa(len(args))
    str += "\r\n"
    var l string
    for _, v := range args  {
        l = strconv.Itoa(len(v))
        str += "$" + l + "\r\n"
        str += v + "\r\n"
    }

    return str
}

func (r *Redis) read() (response string) {
    buffer := make([]byte, 128)
    length, _ := r.conn.Read(buffer)
    return socket_read(string(buffer[:length]))
}

func (r *Redis) write(request string) {
    r.conn.Write([]byte(request))
}

func socket_read(response string) string{
    var res string
    switch (string(response[0])) {
        case "$":
            s := strings.Split(response, "\r\n")
            res = s[1]
            break;
        case "+":
            s := strings.Split(response, "\r\n")
            if string(s[0][1:]) == "OK" {
                res = "OK"
            }
            break;
        case "*":
            break;
    }
    return res
}

func redis_connect(ip, port string) *net.TCPConn {
    tcpAddr, err := net.ResolveTCPAddr("tcp4",  ip + ":" + port)
    if err != nil {
        fmt.Println("error")
        os.Exit(0)
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        fmt.Println("error")
        os.Exit(0)
    }
    return conn
}

func main () {
    redis := New("127.0.0.1", "6379")
    fmt.Println(redis.Set("fuck", "200"))
    fmt.Println(redis.Get("b"))
}
