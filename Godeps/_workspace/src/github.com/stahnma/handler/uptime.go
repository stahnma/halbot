package handler

import (
	"github.com/danryan/hal"
)

type uptime struct {
	hal.Handler
}

func (h *uptime) Method() string {
	return hal.RESPOND
}

func (h *uptime) Usage() string {
	return `uptime - responds with "PONG"`
}

func (h *uptime) Pattern() string {
	return `(?i)uptime`
}

func (h *uptime) Run(res *hal.Response) error {
	return res.Send("uptime")
}

// Uptime exports
var Uptime = &uptime{}
