package main

import (
	"fmt"
	"strconv"
)

type List interface {
    Add(element int)
    Remove(index int) int
    Get(index int) int
    Size() int
    isEmpty() bool
    Contains(element int) bool
}

type ArrayList struct {
    data []int
}

func (a *ArrayList) Add(element int) {
    a.data = append(a.data, element)
    fmt.Println("Element has been added to the ArrayList")
}

func (a *ArrayList) Remove(index int) int {
    if index < 0 || index >= len(a.data) {
        return 0
    }
    removedElement := a.data[index]
    a.data = append(a.data[:index], a.data[index+1:]...)
    fmt.Printf("%s element has been deleted from ArrayList!",strconv.Itoa(a.data[index]))
    fmt.Println()
    return removedElement
}


func (a *ArrayList) Get(index int) int {
    if index < 0 || index >= len(a.data) {
        return 0
    }
    element := strconv.Itoa(a.data[index])
    fmt.Printf("This is your element %s from ArrayList", element)
    return a.data[index]
}

func (a *ArrayList) Size() int {
    size := strconv.Itoa(len(a.data))
    fmt.Printf("Size of your ArrayList: %s", size)
    return len(a.data)
}
func(a *ArrayList) IsEmty() bool{
    if len(a.data) == 0{
        fmt.Println("It is empty!")
        return true;
    }
    fmt.Println("It is not empty!")
    return false;
}
func(a *ArrayList) Contains(element int) bool{
    for _,v := range a.data{
        if v == element{
            fmt.Println("Yes it is found!")
            return true
        }
    }
    return false
}


type Vector struct {
    data []int
}

func (v *Vector) Add(element int) {
    v.data = append(v.data, element)
    fmt.Println("Element has been added to the Vector")
}

func (v *Vector) Remove(index int) int {
    if index < 0 || index >= len(v.data) {
        return 0
    }
    removedElement := v.data[index]
    v.data = append(v.data[:index], v.data[index+1:]...)
    fmt.Printf("%s element has been deleted from Vector!",strconv.Itoa(v.data[index]))
    fmt.Println()
    return removedElement
}

func (v *Vector) Get(index int) int {
    if index < 0 || index >= len(v.data) {
        return 0
    }
    element := strconv.Itoa(v.data[index])
    fmt.Printf("This is your element %s from ArrayList", element)
    return v.data[index]
}

func (v *Vector) push_front(element int) {
    v.data = append([]int{element}, v.data...)
}

func (v *Vector) Size() int {
    size := strconv.Itoa(len(v.data))
    fmt.Printf("Size of your ArrayList: %s", size)
    return len(v.data)
}
func(v *Vector) IsEmty() bool{
    if len(v.data) == 0{
        fmt.Println("It is empty!")
        return true;
    }
    fmt.Println("It is not empty!")
    return false;
}
func(v *Vector) Contains(element int) bool{
    for _,v := range v.data{
        if v == element{
            fmt.Println("Yes it is found!")
            return true
        }
    }
    return false
}
func (v *Vector) pop_back() (int) {
    if len(v.data) == 0 {
        return 0
    }
    lastIdx := len(v.data) - 1
    lastVal := v.data[lastIdx]
    v.data = v.data[:lastIdx]
    return lastVal
}

func (v *Vector) pop_front() (int) {
    if len(v.data) == 0 {
        return 0
    }
    firstVal := v.data[0]
    v.data = v.data[1:]
    return firstVal
}




func main(){
    arr := ArrayList{}
    arr.Add(7)
    arr.Add(9)
    arr.Add(3)
    vec := Vector{}
    vec.Add(6)
    vec.Add(3)
    vec.Add(8)
    arr.Remove(0)
    vec.Remove(0)
    arr.Contains(9)
    vec.Contains(3)
    vec.push_front(8)
    arr.IsEmty()
    vec.IsEmty()
    vec.pop_back()
    vec.pop_front()
    fmt.Println(arr)
    fmt.Println(vec)
}