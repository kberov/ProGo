package main

import (
	"context"
	"math"
	"math/rand"
	"sync"
	"time"
)

// new returns a pointer to the created value
var waitGroupP = new(sync.WaitGroup)
var mutexP = new(sync.Mutex)

/*
It is important not to copy WaitGroup values because it means that goroutines
will be calling Done and Wait on different values, which generally means that
the application deadlocks. As a rule of thumb, coordination requires that all
goroutines use the same struct value.
*/
func doSum(count int, val *int) {
	// The call to the time.Sleep function is ONLY to ensure that the
	// goroutines are all running at once.
	//time.Sleep(time.Second)
	// Perform fewer lock operations. Not on every increment, but once as we
	// start work with the variable.
	mutexP.Lock()
	for i := 0; i < count; i++ {
		*val++
	}
	mutexP.Unlock()
	waitGroupP.Done()
}

var rwmutex = sync.RWMutex{}
var squares = map[int]int{}

func calculateSquares(max, iterations int) {
	for i := 0; i < iterations; i++ {
		val := rand.Intn(max)
		rwmutex.RLock()
		square, ok := squares[val]
		rwmutex.RUnlock()
		if ok {
			Printfln("Cached value: %v = %v", val, square)
		} else {
			rwmutex.Lock()
			if _, ok := squares[val]; !ok {
				squares[val] = int(math.Pow(float64(val), 2))
				Printfln("Added value: %v = %v", val, squares[val])
			}
			rwmutex.Unlock()
		}
	}
	waitGroupP.Done()
}

var readyCond = sync.NewCond(rwmutex.RLocker())

func generateSquares(max int) {
	rwmutex.Lock()
	Printfln("Start Generating data... at %v", time.Now())
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting condition")
	readyCond.Broadcast()
	waitGroupP.Done()
}

func readSquares(id, max, iterations int) {
	Printfln("Start waiting  for RLock at %v", time.Now())
	readyCond.L.Lock()
	for len(squares) == 0 {
		readyCond.Wait()
	}
	Printfln("Starting work after waiting at %v", time.Now())
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v. Sleeping for 100 ms.", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()
	waitGroupP.Done()
}

var once = sync.Once{}

func generateSquaresOnce(max int) {
	// rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	// 826 Chapter 30 ■ Coordinating Goroutines
	// rwmutex.Unlock()
	// Printfln("Broadcasting condition")
	// readyCond.Broadcast()
	// waitGroup.Done()
}
func readSquaresOnce(id, max, iterations int) {
	once.Do(func() {
		generateSquaresOnce(max)
	})
	// readyCond.L.Lock()
	// for len(squares) == 0 {
	//  readyCond.Wait()
	// }
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	// readyCond.L.Unlock()
	waitGroupP.Done()
}

// Simulating processing of a request
func processRequest(wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		Printfln("Processing request: %v", total)
		total++
		time.Sleep(time.Millisecond * 250)
	}
	Printfln("Request processed...%v", total)
	wg.Done()
}

func processRequestOrCancel(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	canceled := false
FOR:
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			Printfln("Stopping processing - request cancelled")
			canceled = true
			break FOR
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	if !canceled {
		Printfln("Request processed...%v", total)
	}
	wg.Done()
}

func processRequestWithDeadline(ctx context.Context, wg *sync.WaitGroup, count int) {
	total := 0
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(time.Millisecond * 250)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}

const (
	countKey = iota
	sleepPeriodKey
)

// 833 Chapter 30 ■ Coordinating Goroutines
func processRequestWithData(ctx context.Context, wg *sync.WaitGroup) {
	total := 0
	count := ctx.Value(countKey).(int)
	sleepPeriod := ctx.Value(sleepPeriodKey).(time.Duration)
	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				Printfln("Stopping processing - request cancelled")
			} else {
				Printfln("Stopping processing - deadline reached")
			}
			goto end
		default:
			Printfln("Processing request: %v", total)
			total++
			time.Sleep(sleepPeriod)
		}
	}
	Printfln("Request processed...%v", total)
end:
	wg.Done()
}
func main() {
	Printfln("\n\nCHAPTER 30\nCoordinating Goroutines\n")

	/*
		Problem | Solution | Listing
		Wait for one or more goroutines to finish | Use a wait group | 5, 6
		Prevent multiple goroutines from accessing data at the same time | Use mutual
		exclusion | 7–10
		Wait for an event to occur | Use a condition | 11, 12
		Ensure a function is executed once | Use a Once struct | 13
		Provide a context for requests being processed across API boundaries in
		servers | Use a context | 14–17
	*/

	counter := 0

	waitGroupP.Add(1)
	go doSum(5000, &counter)
	waitGroupP.Wait()
	Printfln("Total: %v", counter)

	Printfln("%s%s",
		"\n Using Wait Groups",
		"\n Using Mutual Exclusion")

	/* A common problem is ensuring that the main function doesn’t finish
	* before the goroutines it starts are complete, at which point the program
	* terminates. */
	numRoutines := 3
	waitGroupP.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go doSum(5000, &counter)
	}
	waitGroupP.Wait()
	Printfln("Total: %v", counter)

	Printfln("\n  Using a Read-Write Mutex")
	/*
		A Mutex treats all goroutines as being equal and allows only one goroutine to
		acquire the lock. The RWMutex struct is more flexible and supports two
		categories of goroutine: readers and writers. Any number of readers can acquire
		the lock simultaneously, or a single writer can acquire the lock. The idea is
		that readers only care about conflicts with writers and can execute
		concurrently with other readers without difficulty.
	*/

	rand.Seed(time.Now().UnixNano())
	// counter := 0
	numRoutines = 3
	waitGroupP.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go calculateSquares(10, 5)
	}
	waitGroupP.Wait()
	Printfln("Cached values: %v", len(squares))

	Printfln("\n Using Conditions to Coordinate Goroutines")
	numRoutines = 2
	waitGroupP.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}
	waitGroupP.Add(1)
	go generateSquares(10)
	waitGroupP.Wait()
	Printfln("Cached values: %v", len(squares))

	Printfln("\n Ensuring a Function Is Executed Once")

	rand.Seed(time.Now().UnixNano())
	numRoutines = 2
	waitGroupP.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquaresOnce(i, 10, 5)
	}
	// waitGroup.Add(1)
	// go generateSquares(10)
	waitGroupP.Wait()
	Printfln("Cached values: %v", len(squares))

	Printfln("\n Using Contexts")
	waitGroupP.Add(1)
	Printfln("Request dispatched...")
	go processRequest(waitGroupP, 10)
	waitGroupP.Wait()

	Printfln("\n  Canceling a Request")
	waitGroupP.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel := context.WithCancel(context.Background())
	go processRequestOrCancel(ctx, waitGroupP, 10)
	time.Sleep(time.Second)
	Printfln("Canceling request by calling cancel()")
	cancel()
	waitGroupP.Wait()

	Printfln("\n  Setting a Deadline")

	waitGroupP.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	go processRequestWithDeadline(ctx, waitGroupP, 10)
	// time.Sleep(time.Second)
	// Printfln("Canceling request")
	// cancel()
	waitGroupP.Wait()

	Printfln("\n  Providing Request Data")

	waitGroupP.Add(1)
	Printfln("Request dispatched...")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	ctx = context.WithValue(ctx, countKey, 4)
	ctx = context.WithValue(ctx, sleepPeriodKey, time.Millisecond*250)
	go processRequestWithData(ctx, waitGroupP)
	// time.Sleep(time.Second)
	// Printfln("Canceling request")
	// cancel()
	waitGroupP.Wait()
}
