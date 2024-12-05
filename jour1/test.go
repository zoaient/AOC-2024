package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed sample.txt
var imput string

func tri(List []int) []int {
	New_List := []int{}
	for i := 0; i < len(List); i++ {
		indice := smallest(List)
		New_List = append(New_List, List[indice])
		List[indice] = 10000000
	}
	return New_List

}
func smallest(List []int) int {
	indice := 0
	smallest := 1000000000
	for i := 0; i < len(List); i++ {
		if List[i] < smallest {
			smallest = List[i]
			indice = i
		}
	}
	return indice
}

func difference(ListA []int, ListB []int) int {
	ecart := 0
	ListA_triee := tri(ListA)
	ListB_triee := tri(ListB)
	for i := 0; i < len(ListA_triee); i++ {
		ecart += abs(ListA_triee[i] - ListB_triee[i])
	}
	return ecart
}

func difference2(ListA []int, ListB []int) int {
	ecart := 0
	for _, v := range ListA {
		ecart += v * occurences(ListB, v)
	}
	return ecart
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func occurences(List []int, a int) int {
	occurences := 0
	for i := 0; i < len(List); i++ {
		if List[i] == a {
			occurences++
		}

	}
	return occurences
}

func main() {
	ListA := make([]int, 0)
	ListB := make([]int, 0)
	a := strings.Split(imput, "\n")
	a = a[:len(a)-1]
	for _, v := range a {
		b := strings.Split(v, "   ")
		n1, _ := strconv.Atoi(b[0])
		n2, _ := strconv.Atoi(b[1])
		ListA = append(ListA, n1)
		ListB = append(ListB, n2)

	}

	fmt.Println(difference2(ListA, ListB))

}
