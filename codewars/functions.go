package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func FindOdd(seq []int) int {
    res := 0
    for _, x := range seq {
		// If an element appears an even number of times, the XOR operation cancels out, and the corresponding bits in res become 0.
        res ^= x
    }
    return res
}

func findTheDupLetters(s string) (c int) {
	// Initializes an empty map h where the keys are Unicode characters (rune) and the values are integers.
	h := map[rune]int{}
	for _, r := range strings.ToLower(s) {
		// Increments the count of character r in the map h.
		// Checks if the count of character r is equal to 2.
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return
}

// https://www.codewars.com/kata/55c6126177c9441a570000cc/train/go
func OrderWeight(str string) string {
	weights := strings.Split(str, " ")
	specialWeights := []SpecialWeight{};

	for _, weight := range weights {
		weightSum := 0
		for _, num := range weight {
			numAsInt, _ := strconv.Atoi(string(num))
			weightSum += numAsInt
		}
		placeHolder := SpecialWeight{weight, weightSum}
		specialWeights = append(specialWeights, placeHolder)
	}

	fmt.Println(specialWeights)
	sort.Sort(ByValue(specialWeights))
	fmt.Println(specialWeights)

  result := ""
  for _, sWeight := range specialWeights {
    result += sWeight.weight + " "
  }
	return strings.TrimSpace(result)
}

type SpecialWeight struct {
	weight string
	value int
}

func (s SpecialWeight) String() string {
	return fmt.Sprintf("%s : %d\n", s.weight, s.value)
}

type ByValue []SpecialWeight 

func (a ByValue) Len() int 				{ return len(a)}
func (a ByValue) Swap(i, j int) 		{ a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool 	{ 
  if a[i].value == a[j].value {
    return a[i].weight < a[j].weight
  } 
  return a[i].value < a[j].value
}