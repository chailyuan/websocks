package client

import (
	"github.com/chailyuan/websocks/core"
	"github.com/chailyuan/websocks/core/mux"
)

func (client *WebSocksClient) OpenMux() (err error) {
	wsConn, _, err := client.dialer.Dial(client.ServerURL.String(), map[string][]string{
		"WebSocks-Mux": {"v0.15"},
	})

	if err != nil {
		return
	}

	ws := core.NewWebSocket(wsConn, client.Stats)

	muxWS := mux.NewMuxWebSocket(ws)
	client.muxGroup.AddMuxWS(muxWS)
	return
}
