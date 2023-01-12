package main

import "time"

func DoWork(done <-chan interface{}, interval time.Duration) (<-chan interface{}, <-chan time.Time) {
	heart := make(chan interface{})
	results := make(chan time.Time)
	go func() {
		defer close(heart)
		defer close(results)

		pulse := time.Tick(interval)
		workGen := time.Tick(interval * 2)
		sendPulse := func() {
			select {
			case heart <- struct{}{}:
			default:
			}
		}

		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		for {
            select {
		    case <-done:
                return
            case <-pulse:
                sendPulse()
            case r := <- workGen:
                sendResult(r)
            }
        }
	}()
	return heart, results
}
