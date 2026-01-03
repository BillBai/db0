package main

import (
	"fmt"
	"time"
)

// learnSlicesAndArrays demonstrates arrays and slices in Go
func learnSlicesAndArrays() {
	fmt.Println("\n=== Slices and Arrays ===")

	// Arrays have fixed size
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println("Array:", arr)

	// Slices are dynamic
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))

	// Slicing a slice
	subSlice := slice[1:4] // [2, 3, 4]
	fmt.Println("Sub-slice [1:4]:", subSlice)

	// Appending to slice
	slice = append(slice, 6, 7, 8)
	fmt.Println("After append:", slice)

	// Make creates a slice with specified length and capacity
	madeSlice := make([]int, 5, 10)
	fmt.Println("Made slice - len:", len(madeSlice), "cap:", cap(madeSlice))
}

// learnMaps demonstrates map usage in Go
func learnMaps() {
	fmt.Println("\n=== Maps ===")

	// Create a map
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	fmt.Println("Ages map:", ages)
	fmt.Println("Alice's age:", ages["Alice"])

	// Check if key exists
	age, exists := ages["David"]
	if exists {
		fmt.Println("David's age:", age)
	} else {
		fmt.Println("David not found")
	}

	// Delete from map
	delete(ages, "Bob")
	fmt.Println("After deleting Bob:", ages)

	// Map literal
	scores := map[string]int{
		"math":    95,
		"science": 88,
		"history": 92,
	}

	// Iterate over map
	for subject, score := range scores {
		fmt.Printf("%s: %d\n", subject, score)
	}
}

// Shape interface demonstrates interfaces in Go
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func learnInterfaces() {
	fmt.Println("\n=== Interfaces ===")

	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	shapes := []Shape{rect, circle}

	for i, shape := range shapes {
		fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f\n",
			i+1, shape.Area(), shape.Perimeter())
	}
}

// learnGoroutines demonstrates concurrent programming
func learnGoroutines() {
	fmt.Println("\n=== Goroutines and Channels ===")

	// Simple goroutine
	go func() {
		fmt.Println("Running in goroutine")
	}()

	// Channel for communication
	messages := make(chan string)

	go func() {
		messages <- "Hello from goroutine"
	}()

	msg := <-messages
	fmt.Println("Received:", msg)

	// Buffered channel
	nums := make(chan int, 3)
	nums <- 1
	nums <- 2
	nums <- 3

	fmt.Println("Buffered channel:", <-nums, <-nums, <-nums)

	// Channel with multiple workers
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker
	go func() {
		for job := range jobs {
			fmt.Println("Processing job:", job)
			time.Sleep(100 * time.Millisecond)
		}
		done <- true
	}()

	// Send jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)

	<-done
	fmt.Println("All jobs completed")
}

// Counter demonstrates methods on structs
type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) GetCount() int {
	return c.count
}

func learnMethods() {
	fmt.Println("\n=== Methods ===")

	counter := Counter{}
	counter.Increment()
	counter.Increment()
	counter.Increment()

	fmt.Println("Counter value:", counter.GetCount())

	// Pointer receiver vs value receiver
	counter2 := &Counter{}
	counter2.Increment()
	fmt.Println("Counter2 value:", counter2.GetCount())
}

// learnDefer demonstrates defer, panic, and recover
func learnDefer() {
	fmt.Println("\n=== Defer, Panic, Recover ===")

	// Defer executes at the end of function
	defer fmt.Println("This runs last")
	defer fmt.Println("This runs second to last")

	fmt.Println("This runs first")

	// Defer with panic/recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// This would panic, but we recover
	riskyFunction()

	fmt.Println("After risky function")
}

func riskyFunction() {
	// Uncomment to see panic/recover in action
	// panic("Something went wrong!")
	fmt.Println("Risky function executed safely")
}

// learnErrorHandling demonstrates error handling patterns
func learnErrorHandling() {
	fmt.Println("\n=== Error Handling ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// RunExamples runs all example functions
func RunExamples() {
	fmt.Println("=== Running Additional Go Examples ===")
	learnSlicesAndArrays()
	learnMaps()
	learnInterfaces()
	learnMethods()
	learnDefer()
	learnErrorHandling()
	learnGoroutines()
}
