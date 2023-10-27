package main

import (
	"fmt"
	"sort"
	"sync"
)


func main () {
	a := 5
	b := 3
	// comment
	var (
		c int
		e error 
	)
	c, e = add(a, b)
	if e != nil {
		fmt.Println(e)
	}
	var n string = "Hello"
	m := "World"
	fmt.Println(n, m, c)

	// list
	var sl []int
	var li [5]int
	sl = make([]int, 3)
	sl[0] = 1
	sl[1] = 2
	sl[2] = 3
	fmt.Println(sl)
	li[0] = 3
	fmt.Println(li)

	// 2-dimension array
	arr1 := [2][3]int {
		{1, 2, 3},
		{4, 5, 6},
	}
	arr2 := [][]int {
		{1,2,3,4},
		{5,6,7},
		{8,9},
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	// append
	arr3 := append(arr2, []int{10, 11, 12})
	fmt.Println(arr3)
	// for loop
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr2[i]); j++ {
			fmt.Println(arr2[i][j])
		}
	}
	
	// string
	s := "Hello"
	fmt.Println(s)
	// byte
	bb := []byte(s)
	bb[0] = 'h'
	s = string(bb)
	fmt.Println(s)
	// rune: code-point row in unicode
	rs := []rune(s)
	rs[0] = 'は'
	s = string(rs)
	fmt.Println(s)
	// 
	var content = `
	複数行の
	文章からなる
	テストです。
	`
	fmt.Println(content)
	// map
	mm := make(map[string]int)
	mm["hello"] = 1
	mm["OK"] = 2
	fmt.Println(mm)
	// map initialization
	mmm := map[string]int {
		"John": 1,
		"Bob": 2,
		"太郎": 3,
	}
	fmt.Println(mmm)
	// for loop with map
	for k, v := range mmm {
		fmt.Printf("key is %v, value is %d\n", k, v)
	}
	keys := []string{}
	for k := range mmm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i, k := range keys {
		fmt.Printf(
			"%d: key is %v, value is %v\n", 
			i, k, mmm[k])
	}

	// struct
	var user1 User
	user1.Name = "たろう"
	user1.Age = 22
	fmt.Println(user1)
	user1.UpdateName("じろう")
	fmt.Println(user1)
	// constructer
	user2 := *NewUser("David", 35)
	fmt.Println(user2)
	// interface
	dog := Dog{}
	DoSpeak(&dog) // &; pointer of a value
	cat := Cat{}
	DoSpeak(&cat)

	// defer
	defer func() {
		fmt.Println("Defer no-name function")
	}()
	defer fmt.Println("Defer2")
	defer fmt.Println("Defer1")
	fmt.Println("Before Defer")

	// goroutine
	var wg sync.WaitGroup
	var mu sync.Mutex
	num := 0
	for i := 0; i < 1000; i++ {
		v := i
		// fmt.Printf("goroutine %d\n", v)
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Printf("goroutine %d\n", v)
			// Prevent num from being updated at the same time
			mu.Lock()
			num += 1
			mu.Unlock()
		}()
	}
	// Wait until wg counter become 0
	wg.Wait()
	fmt.Printf("Num: %d\n", num)

	// channel
	ch := make(chan string)
	go server(ch)
	for ss := range ch {
		fmt.Println(ss)	
	}
	// var ss string
	// ss = <- ch
	// fmt.Println(ss)
	// ss = <- ch
	// fmt.Println(ss)
	// ss = <- ch
	// fmt.Println(ss)
	ch2 := make(chan string)
	go server(ch2)
	counter := 0
	for {
		select {
		case ss := <- ch2:
			if ss == "" {
				fmt.Println("Completed")
				return
			}else{
				fmt.Println(ss)
			}
		default:
			counter ++
			if counter > 1000 {
				break
			} else {
				fmt.Printf("waiting %d\n", counter)
			}
		}
	}

}

func add (a int, b int) (int, error) {
	c := a + b
	return c, nil
}

type User struct {
	Name string
	Age int
}

// *; type definition for a pointer of a value
func (user *User) UpdateName (newName string) {
	oldName := user.Name
	user.Name = newName
	fmt.Printf("Update Name; %v --> %v\n", oldName, newName)
}

func NewUser(name string, age int) *User {
	return &User {
		Name: name,
		Age: age,
	}
}

type Speaker interface {
	Speak() error
}

type Dog struct {}

func (d *Dog) Speak() error {
	fmt.Println("BowWow")
	return nil
}

type Cat struct {}

func (c *Cat) Speak() error {
	fmt.Println("Meow")
	return nil
}

func DoSpeak (s Speaker) error {
	return s.Speak()
}

// channel
func server (ch chan string) {
	defer close(ch)
	ch <- "one"
	ch <- "two"
	ch <- "three"
}