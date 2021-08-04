package ajson

import (
	"encoding/json"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

// 类型自定义的UnmarshalJSON方法（这里receiver是指针）
// 实现把 a 从 string 转对应的 Animal 类型
func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

// 类型自定义的MarshalJSON方法
// 实现把 a 从 Animal 转对应的 string 类型
func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

//func main() {
//	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
//	var zoo []Animal
//	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("%v\n", zoo)
//
//	census := make(map[Animal]int)
//	for _, animal := range zoo {
//		census[animal] += 1
//	}
//
//	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
//		census[Gopher], census[Zebra], census[Unknown])
//
//	banks := []Animal{
//		0, 1, 2,
//	}
//	enc, err := json.Marshal(banks)
//	fmt.Printf("%s %v\n", enc, err)
//
//	a1 := Animal(1)
//	enc2, err2 := json.Marshal(a1)
//	fmt.Printf("%s %v\n", enc2, err2)
//}
