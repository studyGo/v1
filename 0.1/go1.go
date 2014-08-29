package main
import "fmt"
import "time"
import "math/rand"


func run(name string, dealy time.Duration) {

    t0 := time.Now()
    fmt.Println(name, "start", t0)
    time.Sleep(dealy)

    t1 := time.Now()
    fmt.Println(name, "sleep", t1)
}

func main() {
    rand.Seed(time.Now().Unix())
    var name string

    for i := 0; i < 3; i++ {
        name = fmt.Sprintf("go_%02d", i)
        go run (name, time.Duration(rand.Intn(5)) * time.Second)
    }
    var input string

    fmt.Scanln(&input)
}
