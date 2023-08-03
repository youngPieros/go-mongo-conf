package main

import "github.com/youngPieros/go-mongo-conf/mongoconf"

func test() {
	variables := []*mongoconf.Variable{mongoconf.FromInteger("rank", 2000)}
	panicmode := mongoconf.CreateMongoConf("test-table", 10, variables)
	panicmode.Sync()
	select {}
}
