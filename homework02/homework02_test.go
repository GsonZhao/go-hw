package homework02

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var (
	failedQuestions []string
	totalQuestions  int
	mu              sync.Mutex
)

func recordResult(t *testing.T, name string) {
	mu.Lock()
	defer mu.Unlock()
	totalQuestions++
	if t.Failed() {
		failedQuestions = append(failedQuestions, name)
	}
}

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()

	// Print summary
	if totalQuestions > 0 {
		fmt.Println("\n---------------------------------------------------")
		fmt.Printf("Total Questions: %d\n", totalQuestions)
		fmt.Printf("Passed: %d\n", totalQuestions-len(failedQuestions))
		fmt.Printf("Failed: %d\n", len(failedQuestions))

		score := float64(totalQuestions-len(failedQuestions)) / float64(totalQuestions) * 100
		fmt.Printf("Score: %.2f%%\n", score)

		if len(failedQuestions) > 0 {
			fmt.Println("Failed Questions:")
			for _, q := range failedQuestions {
				fmt.Printf("- %s\n", q)
			}
		}
		fmt.Println("---------------------------------------------------")
	}

	os.Exit(code)
}

func mut_func(p *int) {
	*p += 10
}

func TestMut(t *testing.T) {
	i := 0
	fmt.Println(i)
	mut_func(&i)
	fmt.Println(i)
}

func mut_slice_func1(p *[]int) {
	s := *p
	for i := range s {
		s[i] *= 2
	}
}

func TestPtrSlice(t *testing.T) {
	i := []int{0, 1, 2}
	fmt.Println(i)
	mut_slice_func1(&i)
	fmt.Println(i)
}

func TestGoroutine1(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				fmt.Printf("odd:%d\n", i)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				fmt.Printf("even:%d\n", i)
			}
		}
	}()

	fmt.Println("waiting for finish....")
	wg.Wait()
}

func TestGoroutine2(t *testing.T) {
	type Task func() string

	tasks := make(chan Task, 10)

	var wg sync.WaitGroup

	processMax := 3

	for i := 0; i < processMax; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				now := time.Now()
				res := task()
				fmt.Println(res + " take:" + time.Since(now).String())
			}
		}()
	}

	go func() {
		for i := 0; i < 10; i++ {
			v := i
			tasks <- func() string {
				time.Sleep(1 * time.Second)
				return fmt.Sprintf("task:%d", v)
			}
		}
		close(tasks)
	}()

	wg.Wait()
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct{}

func (Rectangle) Area() {
	fmt.Println("(Rectangle)Area")
}
func (Rectangle) Perimeter() {
	fmt.Println("(Rectangle)Perimeter")
}

type Circle struct{}

func (Circle) Area() {
	fmt.Println("(Circle)Area")
}
func (Circle) Perimeter() {
	fmt.Println("(Circle)Perimeter")
}

func TestOject1(t *testing.T) {
	var s Shape

	circle := Circle{}
	rect := Rectangle{}

	s = circle
	s.Area()
	s.Perimeter()

	s = rect
	s.Area()
	s.Perimeter()
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println(e.Name + ":" + strconv.Itoa(e.Age) + ":" + e.EmployeeID)
}

func TestObject2(t *testing.T) {

	e := Employee{
		Person: Person{Name: "lilei",
			Age: 28},
		EmployeeID: "1000",
	}

	e.PrintInfo()

}

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count += 1
}

func (c *Counter) getCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func TestMutex(t *testing.T) {

	c := Counter{}
	var wg sync.WaitGroup

	processMax := 10
	for i := 0; i < processMax; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				c.increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println(c.getCount())

}

func TestAtomic(t *testing.T) {

	var n atomic.Int32
	var wg sync.WaitGroup
	processMax := 10
	for i := 0; i < processMax; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				n.Add(1)
			}
		}()
	}

	wg.Wait()

	fmt.Println(n.Load())

}
