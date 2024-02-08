package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
)
func main(){
    a1 := []int{1,3}
    a2 := []int{2}

    median := findMedianSortedArrays(a1, a2)
    fmt.Println(median)
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    checkArray1 := checkArray(nums1)
    checkArray2 := checkArray(nums2)

    if !(checkArray1 && checkArray2){
        panic(errors.New("c1"))
    }

    cL := checkLens(len(nums1), len(nums2))
    if !cL {
        panic(errors.New("c2"))
    }

    a := append(nums1, nums2...)
    l := len(a)
    m := l/2

    sort.Slice(a, func(i, j int)bool{
        return a[i]<a[j]
    })

    if l == 1 {
        return float64(a[0])
    }

    if l % 2 != 0 {
        return float64(a[m])
    } else {
        return (float64(a[m]) +  float64(a[m-1]))/2
    }
}

func checkLens(m int, n int) bool {
    s := m + n
    if s > 2000 || s < 1 {
        return false
    }

    if m < 0 || n < 0 {
        return false
    }

    if m > 1000 || n > 1000{
        return false
    }

    return true
}

func checkArray(array []int) bool {
    for _, a := range array{
        if float64(a) < -math.Pow10(6) || float64(a) > math.Pow10(6){
            return false
        }
    }

    return true
}