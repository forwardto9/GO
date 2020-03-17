package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

// Pi const
const Pi = 3.14

// People this is people info
type People struct {
	age  int    //年龄
	name string // 姓名
}

// Car this is a new interface
type Car interface {
	run()
}

// BMW this is a new type of Car
type BMW struct {
}

// Benzi this is a new type of Car
type Benzi struct {
}

func (bmw BMW) run() {
	fmt.Println("Im BMW")
}

func (bmw BMW) String() string {
	return fmt.Sprintf("%s is my description", "BMW")
}

func (benzi Benzi) run() {
	fmt.Println("Im Benzi")
}

func main() {
	/* 这是我的第一个简单的程序 */
	fmt.Println("Hello, World!")
	var sum = sumNumber(1, 2)
	fmt.Println(sum)

	fmt.Println(sumNumberWithError(-1, 0))
	fmt.Println(sumNumberWithError(1, 0))

	var stringArray [2]string
	stringArray[0] = "uwei"
	stringArray[1] = "yuan"
	fmt.Println(stringArray)
	var intArray = [3]int{1, 2, 3}
	fmt.Println(intArray)

	var lily = People{age: 10, name: "lily"}
	fmt.Println(lily.name)
	var tom = People{20, "tom"}
	fmt.Println(tom)

	var slice []float32
	slice = append(slice, 0.1)
	slice = append(slice, 1.2, 3.4, 5.6)
	for _, num := range slice {
		fmt.Println(num)
	}

	var mapInstance map[string]People
	mapInstance = make(map[string]People)
	mapInstance["lily"] = lily
	fmt.Println(mapInstance)

	var car Car = new(BMW)
	car.run()
	fmt.Println(car)
	car = new(Benzi)
	car.run()

	// 在函数中，`:=` 简洁赋值语句在明确类型的地方，可以用于替代 var 定义
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	// 等上层函数执行结束(此处是main)，然后才能执行，是一个压栈操作
	defer fmt.Println("last and end")

	fmt.Println("first")

	pos := adder()
	fmt.Println(pos(1))

	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v.Abs())

	r := strings.NewReader("Hello, Reader!")
	bb := make([]byte, 8)
	for {
		// Read 用数据填充指定的字节 slice，并且返回填充的字节数和错误信息。 在遇到数据流结尾时，返回 io.EOF 错误
		n, err := r.Read(bb)
		fmt.Printf("n = %v err = %v bb = %v\n", n, err, bb)
		fmt.Printf("bb[:n] = %q\n", bb[:n])
		if err == io.EOF {
			break
		}
	}

	channelA := []int{7, 2, 8, -9, 4, 0}

	channelC := make(chan int)
	go summm(channelA[:len(channelA)/2], channelC)
	go summm(channelA[len(channelA)/2:], channelC)

	// 默认情况下，在另一端准备好之前，发送和接收都会阻塞
	firstC, secondC := <-channelC, <-channelC // 从 c 中获取
	fmt.Println(firstC, secondC, firstC+secondC)
	// for value := range channelC {
	// 	fmt.Println(value)
	// }
	// channel 可以用来在多个go routine之间做状态同步，类似 done <- 1(写入), <-done(等待读)
	naturals := make(chan int)
	squares := make(chan int)
	go counter(15, naturals)
	go squarer(naturals, squares)
	printSquarer(squares)

	// go reflect
	reflectInt := 1
	fmt.Println(reflect.TypeOf(reflectInt))
	fmt.Println(reflect.ValueOf(reflectInt).Type().String())
	fmt.Println(reflect.TypeOf(car))
	fmt.Println(reflect.TypeOf(naturals))

	// 类型断言
	var w io.Writer = os.Stdout
	if f, ok := w.(*os.File); ok {
		fmt.Println(ok, f)
		fmt.Fprintf(f, "%d\n", 123)
		w.Write([]byte("fuckyou\n"))
	}

	// Web测试
	// var h Hello
	// err := http.ListenAndServe("localhost:4000", h)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func sumNumber(n1, n2 int) int {
	return n1 + n2
}

func swap(x, y string) (string, string) {
	return y, x
}

// MyError is error resolver
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func sumNumberWithError(x1, y1 int) (int, error) {
	sum := x1 + y1
	if sum < 0 {
		return sum, &MyError{
			time.Now(),
			"it didn't work",
		}
	}
	return sum, nil
}

// 闭包，返回一个返回值是int的有一个参数的匿名函数
func adder() func(int) int {
	sum := 1
	return func(x int) int {
		sum += x
		return sum
	}
}

// Vertex type
type Vertex struct {
	X, Y float64
}

// Abs method, 为结构体定义方法，此种声明可以针对任意类型
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale method
// 当 v 是 Vertex 的时候 Scale 方法没有任何作用，因为是值传递；
// 当是指针的时候，就是引用传递
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Hello is test web
type Hello struct{}

var mu sync.Mutex

// http.Handler interface
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 开启一个新的 goroutine 执行, goroutine 是由 Go 运行时环境管理的轻量级线程
	// goroutine 在相同的地址空间中运行
	mu.Lock()
	// go say("fuck 2 !", w)
	mu.Unlock()
	say("fuck 1 !", w)
}

func say(s string, w http.ResponseWriter) {
	for i := 0; i < 5; i++ {
		// time.Sleep(1000 * time.Millisecond)
		fmt.Fprintf(w, "%q\n", s)
	}
	go fmt.Println("fuck go routine!")
}

func summm(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
	// close(c)
}

// serial go routine
func counter(count int, out chan<- int) {
	for index := 0; index < count; index++ {
		out <- index
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printSquarer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
