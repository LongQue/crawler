package engine

import (
	"awesomeProject1/crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request: %s", r.Url)
		}
		e.Scheduler.Submit(r)
	}
	profileCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			//强转，用ok判断类型，只记录Profile的次数
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got Profile #%d: %v", profileCount, item)
				profileCount++
			}
		}

		//URL dedup
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				log.Printf("Duplicate request: %s", request.Url)
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true
	return false
}
