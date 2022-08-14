package async

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/samber/mo"
	"github.com/stretchr/testify/require"
)

var data = map[string]struct{}{
	"foo": {},
	"bar": {},
}

var notFound = errors.New("not found")

func BlockGetName(name string) (string, error) {
	time.Sleep(1 * time.Second)
	if _, ok := data[name]; !ok {
		return "", notFound
	}

	return name, nil
}

func AsyncGetName(name string) *mo.Future[string] {
	return mo.NewFuture(func(resolve func(string), reject func(error)) {
		name, err := BlockGetName(name)
		if err != nil {
			reject(err)
			return
		}

		resolve(name)
	})
}

func TestFuture(t *testing.T) {
	fuOk := mo.NewFuture(func(resolve func(string), reject func(error)) {
		name, err := BlockGetName("foo")
		if err != nil {
			reject(err)
			return
		}

		resolve(name)
	})

	collect, err := fuOk.Collect()
	require.Nil(t, err)
	require.Equal(t, collect, "foo")
}

func TestFutureFailed(t *testing.T) {
	fuBad := mo.NewFuture(func(resolve func(string), reject func(error)) {
		name, err := BlockGetName("xxx")
		if err != nil {
			reject(err)
			return
		}

		resolve(name)
	})
	collect, err := fuBad.Collect()
	require.Equal(t, err, notFound)
	require.Equal(t, collect, "")
}

func TestAsyncGetName(t *testing.T) {
	fmt.Println(time.Now().Second())
	fooasync := AsyncGetName("foo")
	xxxasync := AsyncGetName("bar")
	fmt.Println("hello world")

	fr := fooasync.Result()
	xR := xxxasync.Result()
	if fr.IsOk() {
		require.Equal(t, "foo", fr.MustGet())
	}
	if xR.IsError() {
		require.Equal(t, notFound, xR.Error())
	}
}
