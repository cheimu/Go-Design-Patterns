package future

import (
	"errors"
	"sync"
	"testing"
)

func TestStringOrError_Execute(t *testing.T) {
	future := &MaybeString{}

	t.Run("Success result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(s string) {
			t.Log(s)
			wg.Done()
		}).Fail(func(_ error) {
			t.Fail()
			wg.Done()

		})
		future.Execute(func() (string, error) {
			return "Hello World!", nil
		})
		wg.Wait()
	})

	t.Run("Error result", func(t *testing.T) {
		var wg sync.WaitGroup
		wg.Add(1)
		future.Success(func(_ string) {
			// should not trigger t.Fail() because it will not success
			t.Fail()
			wg.Done()
		}).Fail(func(e error) {
			t.Log(e.Error())
			wg.Done()
		})
		future.Execute(func() (string, error) {
			return "", errors.New("Error ocurred")
		})
		wg.Wait()
	})
}
