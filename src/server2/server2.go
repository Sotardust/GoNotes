package main

import (
	"github.com/xhyonline/xutil/helper"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		for {
			go func() {
				time.Sleep(time.Second * time.Duration(helper.GetRandom(10)))
			}()
			time.Sleep(time.Second * time.Duration(helper.GetRandom(10)))
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe("0.0.0.0:2113", nil)
}
