package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/hipchat"
	_ "github.com/danryan/hal/store/memory"
	"github.com/stahnma/handler"
	"os"
)

//var echoHandler = hal.Respond(`echo (.+)`, func(res *hal.Response) error {
//  return res.Reply(res.Match[1])
//})

var fooHandler = hal.Respond(`foo`, func(res *hal.Response) error {
	return res.Send("BAR")
})

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	tableFlipHandler := &hal.BasicHandler{
		Method:  hal.HEAR,
		Pattern: `tableflip`,
		Run: func(res *hal.Response) error {
			return res.Send(`(╯°□°）╯︵ ┻━┻`)
		},
	}

	beerHandler := &hal.BasicHandler{
		Method:  hal.HEAR,
		Pattern: `beer`,
		Run: func(res *hal.Response) error {
			return res.Reply(`delicious...beer`)
		},
	}

	robot.Handle(
		//echoHandler,
		//handler.Echo,
		handler.Ping,
		handler.Uptime,
		tableFlipHandler,
		beerHandler,
		fooHandler,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
