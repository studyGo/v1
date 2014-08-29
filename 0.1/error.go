package main

import "fmt"
import "errors"

type myError struct{
    arg int
    errMsg string
}

func (e *myError) Error() string{
    return fmt.Sprintf("%d---%s",e.arg, e.errMsg)
}

func error_test(arg int) (int, error) {

    return -1, &myError{100, "this is a error"}
    return -1, errors.New("this is a error")
}

func main() {
    r, e := error_test(1)
    fmt.Println(r, e)


}

