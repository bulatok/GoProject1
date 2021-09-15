package main

import (
	"fmt"
	//"bufio"
	//"os"
	//"strings"
)

//type Car struct{
//	speed int
//	name string
//	dsl string
//}
//type Moto struct{
//	mv int
//	name string
//	vvs string
//}
//
//type Rider interface {
//	MoveSpeed(int) error
//}
//func (m *Moto) MoveSpeed(int speed) error{
//	m.mv -= speed
//	return nil
//}
//
//func (c *Car) MoveSpeed(int speed) error{
//	c.speed -= speed
//	return nil
//}
func main() {
	const s = 'a'
	fmt.Printf("%T\n", s)
	fmt.Printf("%+q\n", s)
	fmt.Printf("%U\n", s)
}