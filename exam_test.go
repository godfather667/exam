/*
exam_test program contains unit test for the functions in exam.go.

1. TestTriangle
2. TestFive
3. TestEqual
*/

package main

import (
	"container/list"
	"testing"
)

// Test Structure for TestTriangle.
type td struct {
	s1, s2, s3 float64
	r          string
}

// Test Structure for Test Five.
type fv struct {
	listTop int
	listRes int
}

// TestTriangle Test Data.
var tl = []td{
	td{
		s1: 1,
		s2: 1,
		s3: 1,
		r:  "Equilateral",
	},
	td{
		s1: 1,
		s2: 2,
		s3: 1,
		r:  "Isosceles",
	},
	td{
		s1: 1,
		s2: 2,
		s3: 2,
		r:  "Isosceles",
	},
	td{
		s1: 2,
		s2: 2,
		s3: 1,
		r:  "Isosceles",
	},
	td{
		s1: 1,
		s2: 2,
		s3: 3,
		r:  "Scalene",
	},
}

// TestFive Test Data.
var tf = []fv{
	fv{
		listTop: 4,
		listRes: -1,
	},
	fv{
		listTop: 5,
		listRes: 0,
	},
	fv{
		listTop: 6.0,
		listRes: 1.0,
	},
	fv{
		listTop: 17,
		listRes: 12,
	},
}

//
// TestTrianglel Unit Test.
//
func TestTriangle(t *testing.T) {
	for i, _ := range tl {
		r := triangleID(tl[i].s1, tl[i].s2, tl[i].s3)
		if tl[i].r != r {
			t.Error("Expected: ", tl[0].r, " got: ", r)
		}
	}
}

//
// TestFive Unit Test.
//
func TestFive(t *testing.T) {
	for j, _ := range tf {
		l := list.New()
		for i := 0; i < tf[j].listTop; i++ {
			l.PushBack(i)
		}
		res := fiveList(l)
		if tf[j].listRes != res {
			t.Error("Expected: ", tf[j].listRes, " got: ", res)
		}
	}
}

//
// TestEqual Unit Test.
//
func TestEqual(t *testing.T) {

	// Create Test Lists.
	list1 := make([]int, 10)
	for i, _ := range list1 {
		list1[i] = i
	}

	list2 := make([]int, 10)
	for j, _ := range list2 {
		list2[j] = 9 - j
	}

	list3 := make([]int, 10)
	for j, _ := range list3 {
		list3[j] = 9 - j
	}
	// Induce Failure Mode for list3
	list3[5] = 1

	// Test List1 and list2 - Should PASS.
	ok := equalList(list1, list2)
	if ok != true {
		t.Error("Expected: ", true, " got: ", ok)
	}

	// Test List1 and list3 - Should FAIL.
	ok = equalList(list1, list3)
	if ok == true {
		t.Error("Expected: ", false, " got: ", ok)
	}
}
