package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed sample.txt
var imput string

func main() {
	Listfirst := make([]int, 0)
	Listfollow := make([][]int, 0)
	a := strings.Split(imput, "\n")
	a = a[:len(a)-1]
	for _, v := range a {
		b := strings.Split(v, ":")
		n1, _ := strconv.Atoi(b[0])
		Listfirst = append(Listfirst, n1)
		c := strings.Split(b[1], " ")
		auxlist := make([]int, 0)
		for _, v1 := range c {
			n2, _ := strconv.Atoi(v1)

			auxlist = append(auxlist, n2)
		}
		auxlist = auxlist[1:]
		Listfollow = append(Listfollow, auxlist)

	}

	SampleListf := make([]int, 0)
	SampleListf = append(SampleListf, 10)
	SampleLists := make([][]int, 0)
	SampleLists_0 := make([]int, 0)
	SampleLists_0 = append(SampleLists_0, 5)
	SampleLists_0 = append(SampleLists_0, 5)
	SampleValuationArray := make([]int, 0)
	SampleValuationArray = append(SampleValuationArray, 1)
	SampleLists = append(SampleLists, SampleLists_0)

	total := 0
	for i := 0; i < len(Listfirst); i++ {
		if all_valuation(Listfirst, Listfollow, i) {
			total += Listfirst[i]
			fmt.Println(i)
		}
	}
	fmt.Println(total)
}

func valuation(List_f []int, List_s [][]int, indice int, valuation_array []int) bool {
	valuation := List_s[indice][0]
	fmt.Println(valuation_array)
	for i := 0; i < len(valuation_array); i++ {
		if valuation_array[i] == 0 {

			valuation += List_s[indice][i+1]
		} else {
			valuation = valuation * List_s[indice][i+1]
		}
	}
	fmt.Println(valuation)
	if valuation == List_f[indice] {
		return true
	} else {
		return false
	}
}

func permut(arr []int, possiblities []int, length int) [][]int {
	var res [][]int

	if len(arr) == length {
		res = append(res, arr)
		return res
	}

	for i := range possiblities {
		nextPermut := permut(append(arr, possiblities[i]), possiblities, length)
		for k := range nextPermut {
			res = append(res, nextPermut[k])
		}
	}
	return res
}

func all_valuation(List_f []int, List_s [][]int, indice int) bool {
	is_ok := false
	arr := make([]int, 0)
	possiblities := make([]int, 0)
	possiblities = append(possiblities, 0)
	possiblities = append(possiblities, 1)
	permutations := permut(arr, possiblities, len(List_s[indice])-1)
	for i := 0; i < len(permutations); i++ {
		if valuation(List_f, List_s, indice, permutations[i]) {
			is_ok = true
		}
	}
	return is_ok
}
