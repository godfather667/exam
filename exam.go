/*
Program exam provides the answer to three exam problems.

It contains three functions:
	1. triangleID: Determine type of triangle.
	2. fiveList: Find fifth Item from end of a list.
	3. equalList: Compare Two Lists for Equality.

A Unit test is included for each function.
*/
package main

import (
	"container/list"
	"fmt"
)

//
// Determine the type of triangle from a list of side lengths.
//
func triangleID(s1, s2, s3 float64) string {
	if (s1 == s2) && (s1 == s3) {
		return "Equilateral"
	}
	if (s1 == s2) || (s1 == s3) || (s2 == s3) {
		return "Isosceles"
	}
	return "Scalene"
}

//
// Find the fifth item from end of list.
//
func fiveList(l *list.List) int {
	five := make([]*list.Element, 5)
	var i = 0

	// Iterate through a given list.
	for e := l.Front(); e != nil; e = e.Next() {
		five[i%5] = e
		i++
	}

	// Handle lists that have less than five values.
	if i < 5 {
		return -1
	}

	// Display fifth from last value.
	return five[i%5].Value.(int)
}

//
// Determine if two lists are different
//
func equalList(list1, list2 []int) bool {
	var ok bool
	for _, iv := range list1 {
		ok = false
		for _, jv := range list2 {
			if list1[iv] == list2[jv] {
				ok = true
				break
			}
		}
		if !ok {
			return ok
		}
	}
	return true
}

func main() {
	fmt.Println("Exam program")
}
