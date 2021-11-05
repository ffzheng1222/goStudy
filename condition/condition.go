package main

import (
    "fmt"
    "sync"
    "time"
)

type Queue struct {
    queue []string
    cond *sync.Cond
}


func main() {
    var strSlice = []string{"tony01", "tony02", "tony03", "tony04", "tony05"}
    //var mapNum int = len(strSlice)

    q := Queue{
        queue: []string{},
        cond: sync.NewCond(&sync.Mutex{}),
    }

    //sleepWaitCond(strSlice, &q)
    wgCond(strSlice, &q)
}


func sleepWaitCond(strSlice []string, q *Queue)  {
    go func() {
       for _, str := range strSlice {
           q.Enqueue(str)
           time.Sleep(time.Microsecond * 100)
       }
       fmt.Println()
    }()

    for _ = range strSlice {
      time.Sleep(time.Second * 1)
      q.Dequeue()
    }
}


func wgCond(strSlice []string, q *Queue)  {
    wg := sync.WaitGroup{}
    wg.Add(len(strSlice))
    for _, str := range strSlice{
        go func() {
            q.Enqueue(str)
            wg.Done()
        }()
    }
    wg.Wait()

    fmt.Println()

    for _ = range strSlice {
        q.Dequeue()
    }
}


func (q *Queue) Enqueue(item string) {
    q.cond.L.Lock()
    defer q.cond.L.Unlock()

    q.queue = append(q.queue, item)
    fmt.Printf("Enqueue: putting %s to queue, notify all\n", item)
    fmt.Printf("q.queue=%v\n", q.queue)
    q.cond.Broadcast()
}


func (q *Queue) Dequeue() (result string) {
    q.cond.L.Lock()
    defer q.cond.L.Unlock()

    if len(q.queue) == 0 {
        fmt.Println("Dequeue: no data available, wait")
        q.cond.Wait()
    }

    result, q.queue = q.queue[0], q.queue[1:]
    fmt.Printf("Dequeue: result=%s, q.queue=%v\n", result, q.queue)
    return
}