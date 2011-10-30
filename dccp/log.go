// Copyright 2011 GoDCCP Authors. All rights reserved.
// Use of this source code is governed by a 
// license that can be found in the LICENSE file.

package dccp

import (
	"runtime/debug"
)

func (c *Conn) logCatchSeqNo(h *Header, seqNos ...int64) {
	if h == nil { 
		return
	}
	for _, seqNo := range seqNos {
		if h.SeqNo == seqNo {
			c.Logger.Logf("conn", "Catch", h.SeqNo, h.AckNo, "Caught SeqNo=%d: %s\n%s", 
				seqNo, h.String(), string(debug.Stack()))
			break
		}
	}
}

func (c *Conn) logState() {
	c.AssertLocked()
	c.Logger.SetState(c.socket.GetState())
}

func (c *Conn) logReadHeader(h *Header) {
	c.Logger.Logf("conn", "Read", h.SeqNo, h.AckNo, h.String())
}

func (c *Conn) logWriteHeader(h *Header) {
	c.Logger.Logf("conn", "Write", h.SeqNo, h.AckNo, h.String())
}

func (c *Conn) logEvent(s string) {
	c.Logger.Logf("conn", "Event", 0, 0, s)
}

func (c *Conn) logWarn(s string) {
	c.Logger.Logf("conn", "Warn", 0, 0, s)
}