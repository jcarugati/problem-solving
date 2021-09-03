package work
import (
	"fmt"
	"sync"
)
func executeParallel(ch chan<- int, functions ...func() int) {
	var wg sync.WaitGroup
	defer close(ch)
	for _, f := range functions {
		wg.Add(1)
		go func(subF func() int) {
			ch <- subF()
			wg.Done()
		}(f)
	}
	wg.Wait()

}

func exampleFunction(counter int) int {
    sum := 0
    for i := 0; i < counter; i++ {
        sum++
    }
    return sum
}

// ParallelExecution Executes funcs concurrently
func ParallelExecution() {
    expensiveFunction := func() int { return exampleFunction(200000000) }  
    cheapFunction := func() int { return exampleFunction(10000000) }

    ch := make(chan int)
    
    go executeParallel(ch, expensiveFunction, cheapFunction)
    
    for result := range ch {
        fmt.Printf("Result: %d\n", result)
    }
}