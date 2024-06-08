package main

import (
	"fmt"

	table "github.com/danipaniii/grafana-template-builder/pkg/panels"
)

func main() {
	fmt.Println("Hello Package")

	a := table.Table

	fmt.Println(a)
}
