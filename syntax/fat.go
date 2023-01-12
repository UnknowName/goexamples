package main

// 重复函数，将value通过通道返回

func Repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	channel := make(chan interface{})
	go func() {
		defer close(channel)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case channel <- v:
				}
			}
		}
	}()
	return channel
}

// 从给出的channel中取出第i个元素，然后退出

func Take(done <-chan interface{}, valueChannel <-chan interface{}, i int) <-chan interface{} {
	channel := make(chan interface{})
	go func() {
		for num := i; num > 0; num-- {
			select {
			case <-done:
				return
			// 从valueChannel取出来之后，再放入channel
			case channel <- <-valueChannel:
			}
		}
	}()
	return channel
}
