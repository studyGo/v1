package main

/*
#include <stdio.h>

void help() {
    printf("this is a demo\n");
    return;
}
*/
import "C"

func main() {
	C.help()
}
