package test

import (
	"log"
	"testing"

	"github.com/wangbaojin/go-okx/ws"
	"github.com/wangbaojin/go-okx/ws/public"
)

func TestBooks1(t *testing.T) {
	args := &ws.Args{
		Channel: "books",
		InstId:  "APT-USDC",
	}
	handler := func(c interface{}) {
		log.Println(c.(public.EventBooks))
	}
	handlerError := func(err error) {
		panic(err)
	}
	if _, err := public.SubscribeBooks(args, handler, handlerError, false); err != nil {
		panic(err)
	}
	select {} // Wait forever
}
