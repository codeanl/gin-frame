package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"rookieCode/routes"
)

var g errgroup.Group

func main() {
	routes.InitGlobalVariable()
	// 后台接口服务

	g.Go(func() error {
		return routes.AdminServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
