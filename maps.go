package main

import "fmt"





func printMaps(m map[int]string)  {
	m[2] = "my name"
	fmt.Println(m)
}

func iterateMaps(m map[int]string)  {
	for _, name := range m {
		fmt.Println(name)
	}
}

