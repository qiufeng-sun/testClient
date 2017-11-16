package main

import (
	"time"

	proto "github.com/golang/protobuf/proto"

	"util/logs"

	. "share/msg"
)

//
var _ = logs.Debug
var _ = time.Sleep
var _ = proto.Bool

//
func TestClient(c *Client) {
	testLogin(c)
	testEnterWorld(c)
}

//
func testLogin(c *Client) {
	var e error

	//
	logon := &CSLogin{}
	e = c.Send(EMsg_ID_CSLogin, logon)
	c.CheckPanic(e)

	var resp SCLogin
	e = c.Recv(EMsg_ID_SCLogin, &resp)
	c.CheckPanic(e)

	logs.Info("%v:%#v", c, resp)
}

//
func testEnterWorld(c *Client) {
	var e error

	m := &CSEnterWorld{}
	e = c.Send(EMsg_ID_CSEnterWorld, m)
	c.CheckPanic(e)

	var resp SCEnterWorld
	e = c.Recv(EMsg_ID_SCEnterWorld, &resp)
	c.CheckPanic(e)

	logs.Info("%v:%#v", c, resp)
}
