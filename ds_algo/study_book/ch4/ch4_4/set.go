package main

import (
	"fmt"

	"github.com/golang-collections/collections/set"
)

func main() {
	st := set.New()
	st.Insert(1)
	st.Insert(2)
	st.Insert(1)

	fmt.Println(st.Has(1))
	fmt.Println(st.Has(3))
	fmt.Println(st.Len())
	fmt.Println(st)
}
