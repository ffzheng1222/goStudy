package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var strSlice = []string{"tony01", "tony02", "tony03", "tony04", "tony05"}
	var mapNum int = len(strSlice)

	unsafeWrite(mapNum, strSlice)
	unsafeWriteBuf(mapNum, strSlice)
	safeWrite(mapNum, strSlice)
	safeWriteBuf(mapNum, strSlice)
	time.Sleep(time.Second)
}

/* unsafeWrite使用不带缓冲区 chan 打印显示写入map的数据 */
func unsafeWrite(mapNum int, strSlice []string) {
	configMap := make(map[string]int, mapNum)
	unswchin := make(chan int)

	//for i := 0; i < mapNum; i++ {
	//	go func(i int) {
	//		kStr := strSlice[i]
	//		configMap[kStr] = i
	//		unswchin <- i
	//	}(i)
	//
	//	if num, ok := <- unswchin; ok{
	//		fmt.Printf("unsafeWrite: num=%d,  configMap=%v\n", num, configMap)
	//	}
	//}

	go func() {
		for i := 0; i < mapNum; i++ {
			kStr := strSlice[i]
			configMap[kStr] = i
			unswchin <- i
			time.Sleep(time.Microsecond * 100)
		}
	}()

	for i := 0; i < mapNum; i++ {
		if num, ok := <-unswchin; ok {
			fmt.Printf("unsafeWrite: num=%d,  configMap=%v\n", num, configMap)
		}
	}
}

/* unsafeWriteBuf使用带缓冲区 chan 打印显示写入map的数据 */
func unsafeWriteBuf(mapNum int, strSlice []string) {
	configMap := make(map[string]int, mapNum)
	unswchin := make(chan int, mapNum)

	for i := 0; i < mapNum; i++ {
		go func(i int) {
			unswchin <- i
		}(i)
		time.Sleep(time.Microsecond * 100)
	}
	for i := 0; i < mapNum; i++ {
		num := <-unswchin
		kStr := strSlice[num]
		configMap[kStr] = num
		fmt.Printf("unsafeWriteBuf: num=%d,  configMap=%v\n", num, configMap)
	}
	defer close(unswchin)

	//go func() {
	//	for i := 0; i < mapNum; i++ {
	//		unswchin <- i
	//	}
	//	defer close(unswchin)
	//}()
	//
	//for num := range unswchin {
	//	kStr := strSlice[num]
	//	configMap[kStr] = num
	//	fmt.Printf("unsafeWriteBuf: num=%d,  configMap=%v\n", num, configMap)
	//}
}

type SafeMap struct {
	safeMap map[string]int
	sync.Mutex
}

/* safeWrite使用不带缓冲区 chan 打印显示写入map的数据 */
func safeWrite(mapNum int, strSlice []string) {
	sm := SafeMap{
		safeMap: map[string]int{},
		Mutex:   sync.Mutex{},
	}

	swchin := make(chan int)
	for i := 0; i < mapNum; i++ {
		go func() {
			kStr := strSlice[i]
			sm.Write(kStr, i)
			swchin <- i
		}()

		if num, ok := <-swchin; ok {
			fmt.Printf("safeWrite: num=%d,  configMap=%v\n", num, sm.safeMap)
		}
	}
}

/* safeWriteBuf使用带缓冲区 chan 打印显示写入map的数据 */
func safeWriteBuf(mapNum int, strSlice []string) {
	sm := SafeMap{
		safeMap: map[string]int{},
		Mutex:   sync.Mutex{},
	}

	swchin := make(chan int, mapNum)
	go func() {
		for i := 0; i < mapNum; i++ {
			swchin <- i
		}
		defer close(swchin)
	}()
	for num := range swchin {
		kStr := strSlice[num]
		sm.safeMap[kStr] = num
		fmt.Printf("safeWriteBuf: num=%d,  configMap=%v\n", num, sm.safeMap)
	}
}

func (s *SafeMap) Write(k string, v int) {
	s.Lock()
	defer s.Unlock()
	s.safeMap[k] = v
}

func (s *SafeMap) Read(k string, v int) (int, bool) {
	s.Lock()
	defer s.Unlock()
	result, ok := s.safeMap[k]
	return result, ok
}
