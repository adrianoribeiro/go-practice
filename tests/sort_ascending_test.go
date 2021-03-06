package main

import (
	"testing"
)

var emptyArray = make([]int, 0)
func SortAscending(items []int) []int {

	if len(items) == 0 {
		return emptyArray
	}

	cloneItems := cloneValues(&items)
	for i:=0;i<len(items)-1;i++ {

		if cloneItems[i] > cloneItems[i+1] {
			aux := cloneItems[i]
			cloneItems[i] = cloneItems[i+1]
			cloneItems[i+1] = aux

			//To restart
			i = -1
		}
	}
	return cloneItems
}

func TestSortAscendingEmpty(t *testing.T) {
	emptyResult := SortAscending(emptyArray)
	if len(emptyResult) != 0{
		t.Error("An empty array must result an empty array as result.")
	}
}

func TestSortAscendingWithUnsortedValuesShouldReturnASortedArray(t *testing.T) {
	unsortedArray := []int{9,8,7,6,5,4,3,2,1}
	sortedArray := []int{1,2,3,4,5,6,7,8,9}

	result := SortAscending(unsortedArray)

	//We can use !reflect.DeepEqual(result, sortedArray){
	if !areEquals(result, sortedArray){
		t.Error("An empty array must result an empty array as result.")
	}
}

func TestSortAscendingWithSortedValuesShouldReturnTheSameArray(t *testing.T) {
	sortedArray := []int{1,2,3,4,5,6,7,8,9}

	result := SortAscending(sortedArray)

	//We can use !reflect.DeepEqual(result, sortedArray){
	if !areEquals(result, sortedArray){
		t.Error("An empty array must result an empty array as result.")
	}
}

//Aux functions
func areEquals(arrayResulted []int, arrayExpected []int) bool {
	if len(arrayResulted) != len(arrayExpected){
		return false
	}

	for i:=0;i<len(arrayResulted)-1;i++ {
		if arrayResulted[i] != arrayExpected[i] {
			return false
		}
	}
	return true
}

func cloneValues(items *[]int) []int {
	cloneItems := make([]int, 0)
	for _, value := range *items {
		cloneItems = append(cloneItems, value)
	}
	return cloneItems
}
