package main

import (
	"fmt"
	"image/color"
)

type picture struct {
	imageOf [5][5]color.RGBA
}

func newPicture() *picture {

}

type drawable interface {
    render() picture
}


type particle struct {
	
}

func (r particle) render() picture {
	rendering := new picture
	return 
}


type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println("Hello, 世界")

	rgba := color.RGBA{R: 1, G: 2, B: 3, A: 0xff}

	fmt.Println(rgba)
	fmt.Println(person{"Bob", 20})
	fmt.Println(newPerson("Jon"))

}

