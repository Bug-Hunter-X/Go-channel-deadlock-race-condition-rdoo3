```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        
        ch := make(chan int)
        
        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := 0; i < 10; i++ {
                        ch <- i
                }
                close(ch) // Close channel explicitly after sending all values
        }()

        wg.Add(1)
        go func() {
                defer wg.Done()
                for i := range ch {
                        fmt.Println(i)
                }
        }()

        wg.Wait()
} 
```