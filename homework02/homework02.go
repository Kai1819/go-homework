package homework02

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

// ## :white_check_mark:指针
// 1. 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func CheckPoint1(num *int) {
	*num += 10
}

// 2. 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func CheckPoint2(slice *[]int) {
	for index, value := range *slice {
		(*slice)[index] = value * 2
	}
}

// ## :white_check_mark:Goroutine
// 1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func CheckGoroutine1() {
	var wg sync.WaitGroup
	wg.Add(2)
	go odd(&wg)
	go even(&wg)
	wg.Wait()

}
func odd(wg *(sync.WaitGroup)) {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		fmt.Println("奇数:", i)
	}
}
func even(wg *(sync.WaitGroup)) {
	defer wg.Done()
	for i := 2; i < 10; i += 2 {
		fmt.Println("偶数:", i)
	}
}

// 2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func CheckGoroutine2() {
	tasks := []task{
		{
			name : "t1",
			job : func(){
				fmt.Println("任务1执行中......")
				time.Sleep(1 * time.Second)
			},
		},
		{
			name : "t2",
			job : func(){
				fmt.Println("任务2执行中......")
				time.Sleep(2 * time.Second)
			},
		},
		{
			name : "t3",
			job : func(){
				fmt.Println("任务3执行中......")
				time.Sleep(3 * time.Second)
			},
		},
	}

	task_scheduler(tasks)

}

type task struct {
	name string
	job func()
}

func task_scheduler(tasks []task) {
	var wg sync.WaitGroup
	for _,tk  := range tasks {
		wg.Add(1)
		go func(t task){
			defer wg.Done()
			start := time.Now()
			t.job()
			duration := time.Since(start)
			fmt.Printf("任务%s：执行完成，耗时: %v\n", t.name, duration)
		}(tk)
	}
	wg.Wait()
}

// ## :white_check_mark:面向对象
// 1. 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
//    - 考察点 ：接口的定义与实现、面向对象编程风格。

func CheckStruct1(){
	shapes := []Shape {
		Rectangle{Name: "矩形", Wide: 4, High: 5},
		Circle{Name: "圆", Radius: 5},
	}
	for _,shape  := range shapes {
		// fmt.Printf("面积：%f，周长：%f\n", shape.Area(), shape.Perimeter())
		PrintShapeInfo(shape)
	}
}

type Shape interface {
	Desc() string
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	Name string
	Wide,High float64
}
func (r Rectangle) Desc() string {
	return r.Name
}
func (r Rectangle) Area() float64 {
	return r.Wide * r.High
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Wide + r.High)
}
type Circle struct {
	Name string
	Radius float64
}
func (c Circle) Desc() string {
	return c.Name
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(shape Shape){
	fmt.Printf("名称：%s：面积：%f，周长：%f\n", shape.Desc(), shape.Area(), shape.Perimeter())
}

// 2. 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//    - 考察点 ：组合的使用、方法接收者。
func CheckStruct2(){
	employee := Employee{ 
		EmployeeID:1, 
		Person: Person{Name:"张三", Age:18},
	}
	PrintInfo(employee)
}
type Person struct {
	Name string
	Age int
}
type Employee struct {
	EmployeeID int 
	Person
}

func PrintInfo(employee Employee) {
	fmt.Printf("ID:%d，姓名：%s，年龄：%d",employee.EmployeeID,employee.Name,employee.Age)
}

// ## :white_check_mark:Channel
// 1. 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
//    - 考察点 ：通道的基本使用、协程间通信。
func CheckChannel1()  {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(2)
	go func(){
		defer wg.Done()
		for v := range ch {
			time.Sleep(50 * time.Millisecond)
			fmt.Println(v)
		}
	}()
	go func(){
		defer wg.Done()
		defer close(ch) // 生产者关闭channel
		for i := 1; i <= 10; i++ {
			time.Sleep(100 * time.Millisecond)
			ch <- i
		}
	}()
	wg.Wait()

}
// 2. 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//    - 考察点 ：通道的缓冲机制。
func CheckChannel2() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(2)
	go func(){
		defer wg.Done()
		for v := range ch {
			time.Sleep(5 * time.Millisecond)
			fmt.Println(v)
		}
	}()
	go func(){
		defer wg.Done()
		defer close(ch) // 生产者关闭channel
		for i := 1; i <= 100; i++ {
			time.Sleep(10 * time.Millisecond)
			ch <- i
		}
	}()
	
	wg.Wait()
}

// ## :white_check_mark:锁机制
// 1. 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//    - 考察点 ： sync.Mutex 的使用、并发数据安全。
func CheckLock1(){
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)
	var num int64 = 0
	for i := 1; i <= 10; i++ {
		go func(){
			defer wg.Done()
			for j := 1; j <= 1000; j++ {
				mu.Lock()
				num ++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}

// 2. 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//    - 考察点 ：原子操作、并发数据安全。
func CheckLock2(){
	var wg sync.WaitGroup
	wg.Add(10)
	var num int64 = 0
	for i := 1; i <= 10; i++ {
		go func(){
			defer wg.Done()
			for j := 1; j <= 1000; j++ {
				atomic.AddInt64(&num, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}