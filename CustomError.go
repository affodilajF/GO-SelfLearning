package main

//
//import (
//	"errors"
//	"fmt"
//)
//
//func main() {
//	fmt.Println(checkYA())
//	fmt.Println()
//}
//
//func checkYA() error {
//	return &NotFoundError{"ini error ya"}
//	//return NotFoundErrorNih
//}
//
//type NotFoundError struct {
//	Message string
//}
//
//func (n *NotFoundError) Error() string {
//	return n.Message
//}
//
//
//// --------------------------------------------------
//// ini kelamaan si kalo pake yang diatas, pake yg dibawah ini aja
//// bisa pake variabel aja
//var (
//	NotFoundErrorNih   = errors.New("Ini error not found")
//	ValidationErrorNih = errors.New("Ini error validation")
//)
