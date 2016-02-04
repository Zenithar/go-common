package eventbus

import (
	"testing"
	"time"
)

func TestLocalNew(t *testing.T) {
	bus := NewLocal()
	if bus == nil {
		t.Log("New EventBus not created!")
		t.Fail()
	}
}

func TestLocalHasCallback(t *testing.T) {
	bus := NewLocal()
	bus.Subscribe("topic", func() {})
	if bus.HasCallback("topic_topic") {
		t.Fail()
	}
	if !bus.HasCallback("topic") {
		t.Fail()
	}
}

func TestLocalSubscribe(t *testing.T) {
	bus := NewLocal()
	if bus.Subscribe("topic", func() {}) != nil {
		t.Fail()
	}
	if bus.Subscribe("topic", "String") == nil {
		t.Fail()
	}
}

func TestLocalSubscribeOnce(t *testing.T) {
	bus := NewLocal()
	if bus.SubscribeOnce("topic", func() {}) != nil {
		t.Fail()
	}
	if bus.SubscribeOnce("topic", "String") == nil {
		t.Fail()
	}
}

func TestLocalUnsubscribe(t *testing.T) {
	bus := NewLocal()
	handler := func() {}
	bus.Subscribe("topic", handler)
	if bus.Unsubscribe("topic", handler) != nil {
		t.Fail()
	}
	if bus.Unsubscribe("topic", handler) == nil {
		t.Fail()
	}
}

func TestLocalPublish(t *testing.T) {
	bus := NewLocal()
	bus.Subscribe("topic", func(a int, b int) {
		if a != b {
			t.Fail()
		}
	})
	bus.Publish("topic", 10, 10)
}

func TestLocalSubcribeOnceAsync(t *testing.T) {
	var results []int

	bus := NewLocal()
	bus.SubscribeOnceAsync("topic", func(a int, out *[]int) {
		*out = append(*out, a)
	})

	bus.Publish("topic", 10, &results)
	bus.Publish("topic", 10, &results)

	bus.WaitAsync()

	if len(results) != 1 {
		t.Fail()
	}

	if bus.HasCallback("topic") {
		t.Fail()
	}
}

func TestLocalSubscribeAsyncTransactional(t *testing.T) {
	var results []int

	bus := NewLocal()
	bus.SubscribeAsync("topic", func(a int, out *[]int, dur string) {
		sleep, _ := time.ParseDuration(dur)
		time.Sleep(sleep)
		*out = append(*out, a)
	}, true)

	bus.Publish("topic", 1, &results, "1s")
	bus.Publish("topic", 2, &results, "0s")

	bus.WaitAsync()

	if len(results) != 2 {
		t.Fail()
	}

	if results[0] != 1 || results[1] != 2 {
		t.Fail()
	}
}

func TestLocalSubscribeAsync(t *testing.T) {
	results := make(chan int)

	bus := NewLocal()
	bus.SubscribeAsync("topic", func(a int, out chan<- int) {
		out <- a
	}, false)

	bus.Publish("topic", 1, results)
	bus.Publish("topic", 2, results)

	numResults := 0

	go func() {
		for _ = range results {
			numResults++
		}
	}()

	bus.WaitAsync()

	if numResults != 2 {
		t.Fail()
	}
}
