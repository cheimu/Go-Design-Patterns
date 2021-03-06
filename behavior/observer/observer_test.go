package observer

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID int
	EVT
}

func (p *TestObserver) Notify(event EVT) {
	fmt.Printf("Observer %d: message '%s' received \n", p.ID, event)
	p.EVT = event
}

func TestSubject(t *testing.T) {
	testObserver1 := &TestObserver{ID: 1, EVT: &Event{Message: "default"}}
	testObserver2 := &TestObserver{ID: 2, EVT: &Event{Message: "default"}}
	testObserver3 := &TestObserver{ID: 3, EVT: &Event{Message: "default"}}

	publisher := Publisher{}

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(testObserver1)
		publisher.AddObserver(testObserver2)
		publisher.AddObserver(testObserver3)

		if len(publisher.ObserversList) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(testObserver2)

		if len(publisher.ObserversList) != 2 {
			t.Errorf("The size of the observer list is not the "+
				"expected. 3 != %d\n", len(publisher.ObserversList))
		}

		for _, observer := range publisher.ObserversList {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
			}

			if testObserver.ID == 2 {
				t.Fail()
			}
		}

	})

	t.Run("Notify", func(t *testing.T) {

		if len(publisher.ObserversList) == 0 {
			t.Errorf("The list is empty. Nothing to test\n")
		}

		for _, observer := range publisher.ObserversList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			event, _ := (printObserver.EVT).(*Event)
			if event.Message != "default" {
				t.Errorf("The observer's Message field weren't"+
					" empty: %s\n", event.Message)
			}
		}

		message := "Hello World!"
		publisher.NotifyObservers(&Event{message})

		for _, observer := range publisher.ObserversList {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			event, _ := (printObserver.EVT).(*Event)
			if event.Message != message {
				t.Errorf("Expected message on observer %d was "+
					"not expected: '%s' != '%s'\n", printObserver.ID,
					event.Message, message)
			}
		}
	})
}
