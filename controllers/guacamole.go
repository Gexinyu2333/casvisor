// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"net/http"
	"strconv"

	"github.com/beego/beego"
	"github.com/casbin/casvisor/object"
	"github.com/casbin/casvisor/util"
	"github.com/casbin/casvisor/util/tunnel"
	"github.com/gorilla/websocket"
)

const (
	TunnelClosed             int = -1
	Normal                   int = 0
	NotFoundSession          int = 800
	NewTunnelError           int = 801
	ForcedDisconnect         int = 802
	AccessGatewayUnAvailable int = 803
	AccessGatewayCreateError int = 804
	AssetNotActive           int = 805
	NewSshClientError        int = 806
)

var UpGrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols: []string{"guacamole"},
}

func (c *ApiController) GetAssetTunnel() error {
	ctx := c.Ctx
	ws, err := UpGrader.Upgrade(ctx.ResponseWriter, ctx.Request, nil)
	if err != nil {
		beego.Error("WebSocket upgrade failed:", err)
		return err
	}

	owner := c.Input().Get("owner")
	name := c.Input().Get("name")
	width := c.Input().Get("width")
	height := c.Input().Get("height")
	dpi := c.Input().Get("dpi")

	asset, err := object.GetAsset(util.GetIdFromOwnerAndName(owner, name))
	if err != nil {
		return err
	}

	configuration := tunnel.NewConfiguration()

	configuration.SetParameter("width", width)
	configuration.SetParameter("height", height)
	configuration.SetParameter("dpi", dpi)

	configuration.Protocol = "rdp"
	configuration.SetParameter("hostname", asset.Ip)
	configuration.SetParameter("port", strconv.Itoa(asset.Port))
	configuration.SetParameter("username", asset.Username)
	configuration.SetParameter("password", asset.Password)
	configuration.SetParameter("security", "any")
	configuration.SetParameter("ignore-cert", "true")
	configuration.SetParameter("resize-method", "reconnect")

	addr := beego.AppConfig.String("guacamoleEndpoint")
	// fmt.Sprintf("%s:%s", configuration.GetParameter("hostname"), configuration.GetParameter("port"))
	// log.Debug("Intializing guacd.go session", log.String("sessionId", sessionId), log.String("addr", addr), log.String("asset", asset))

	guacdTunnel, err := tunnel.NewTunnel(addr, configuration)
	if err != nil {
		tunnel.Disconnect(ws, NewTunnelError, err.Error())
		// log.Error("Failed to start session", log.String("sessionId", sessionId), log.NamedError("err", err))
		panic(err)
	}

	guacamoleHandler := NewGuacamoleHandler(ws, guacdTunnel)
	guacamoleHandler.Start()
	defer guacamoleHandler.Stop()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			// log.Debug("WebSocket shutdown", log.String("sessionId", sessionId), log.NamedError("err", err))
			// guacdTunnel.Read()
			_ = guacdTunnel.Close()

			return nil
		}
		_, err = guacdTunnel.WriteAndFlush(message)
		if err != nil {
			//service.SessionService.CloseSessionById(sessionId, TunnelClosed, "Remote connection shut down")
			//panic(err)
			return nil
		}
	}
}
