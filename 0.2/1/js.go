package main

import "fmt"
import "encoding/json"
import "strings"

type Data struct {
	data1 string
	data2 string
}

type Action struct {
	T string
}

func main() {
	w := "123t,3465"
	wws := strings.Split(w, ",")
	cc(wws)

	s := `{"t":"sdsd"}`
	dec := json.NewDecoder(strings.NewReader(s))
	var a Action
	dec.Decode(&a)
	fmt.Println(a.T)

}
func cc(s []string) {
	fmt.Println(s)
}
