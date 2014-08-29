package main
import "fmt"

func f(msg string) {
    fmt.Println(msg)
}

func main() {
    go f("this is go demo")
    go func (msg string) {
        fmt.Println(msg)
    }("going")

    // 必须让主线程等待 主线程推出 自线程也退出  就不会输出上面的 字符了
    var input string
    fmt.Scanln(&input)
}
