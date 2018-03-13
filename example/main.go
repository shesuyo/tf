package main

import (
	"fmt"

	"github.com/shesuyo/wtf"
)

func main() {
	{
		a := []string{"1", "2", "who"}
		fmt.Println(wtf.Strings2Interfaces(a), len(wtf.Strings2Interfaces(a)), cap(wtf.Strings2Interfaces(a)))
	}
}
