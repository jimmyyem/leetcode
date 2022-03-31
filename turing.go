package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width, height float64
}

func (c *Rect) Area() float64 {
	return c.width * c.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {

	r := Rect{2, 3}
	var s Shape = r
	area := s.Area()
	fmt.Println(area)

	//s2 := []int{1, 2, 3}
	//s3 := []int{4, 5, 6, 7}
	//n2 := copy(s2, s3)
	//fmt.Printf("n2=%d s2=%v, s3=%v\n", n2, s2, s3)

	//m := make(map[int]int)
	//wg := &sync.WaitGroup{}
	//mu := &sync.Mutex{}
	//wg.Add(10)
	//
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		defer wg.Done()
	//		mu.Lock()
	//		fmt.Println(i)
	//		m[i] = i
	//		mu.Unlock()
	//	}()
	//}
	//
	//wg.Wait()
	//fmt.Println(len(m))
	//fmt.Printf("%+v\n", m)

}
