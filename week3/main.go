package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"net/http"
)

func serverStart() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(resp, "start http server")
	})
	if err := http.ListenAndServe(":9527", mux); err != nil {
		errors.WithMessage(err, "初始化http失败")
		return err
	}
	return nil
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	group := new(errgroup.Group)
	group.Go(func() error {
		done <- serverStart(stop)

		return nil
	})

	if err := group.Wait(); err != nil {

	}
}

func t1() {

	// 声明一个 errgroup 的变量
	group := new(errgroup.Group)
	nums := []int{-1, 0, 1, 2, 3}

	for _, num := range nums {
		midNum := num
		// 开启子协程
		group.Go(func() error {
			if midNum < 1 {
				return errors.New("midNum < 0")
			}
			fmt.Printf("midNum is %v\n", midNum)
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println("catch the err:", err)
		return
	}
	fmt.Println("All is well")
}
