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

package routers

import (
	"github.com/beego/beego"
	"github.com/casvisor/casvisor/controllers"
)

func init() {
	initAPI()
}

func initAPI() {
	ns := beego.NewNamespace("/api",
		beego.NSInclude(
			&controllers.ApiController{},
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/api/signin", &controllers.ApiController{}, "POST:Signin")
	beego.Router("/api/signout", &controllers.ApiController{}, "POST:Signout")
	beego.Router("/api/get-account", &controllers.ApiController{}, "GET:GetAccount")

	beego.Router("/api/get-records", &controllers.ApiController{}, "GET:GetRecords")
	beego.Router("/api/get-record", &controllers.ApiController{}, "GET:GetRecord")
	beego.Router("/api/update-record", &controllers.ApiController{}, "POST:UpdateRecord")
	beego.Router("/api/add-record", &controllers.ApiController{}, "POST:AddRecord")
	beego.Router("/api/delete-record", &controllers.ApiController{}, "POST:DeleteRecord")

	beego.Router("/api/commit-record", &controllers.ApiController{}, "POST:CommitRecord")
	beego.Router("/api/query-record", &controllers.ApiController{}, "GET:QueryRecord")

	beego.Router("/api/get-assets", &controllers.ApiController{}, "GET:GetAssets")
	beego.Router("/api/get-asset", &controllers.ApiController{}, "GET:GetAsset")
	beego.Router("/api/update-asset", &controllers.ApiController{}, "POST:UpdateAsset")
	beego.Router("/api/add-asset", &controllers.ApiController{}, "POST:AddAsset")
	beego.Router("/api/delete-asset", &controllers.ApiController{}, "POST:DeleteAsset")

	beego.Router("/api/get-providers", &controllers.ApiController{}, "GET:GetProviders")
	beego.Router("/api/get-provider", &controllers.ApiController{}, "GET:GetProvider")
	beego.Router("/api/update-provider", &controllers.ApiController{}, "POST:UpdateProvider")
	beego.Router("/api/add-provider", &controllers.ApiController{}, "POST:AddProvider")
	beego.Router("/api/delete-provider", &controllers.ApiController{}, "POST:DeleteProvider")

	beego.Router("/api/get-machines", &controllers.ApiController{}, "GET:GetMachines")
	beego.Router("/api/get-machine", &controllers.ApiController{}, "GET:GetMachine")
	beego.Router("/api/update-machine", &controllers.ApiController{}, "POST:UpdateMachine")
	beego.Router("/api/add-machine", &controllers.ApiController{}, "POST:AddMachine")
	beego.Router("/api/delete-machine", &controllers.ApiController{}, "POST:DeleteMachine")

	beego.Router("/api/get-sessions", &controllers.ApiController{}, "GET:GetSessions")
	beego.Router("/api/get-session", &controllers.ApiController{}, "GET:GetConnSession")
	beego.Router("/api/update-session", &controllers.ApiController{}, "POST:UpdateSession")
	beego.Router("/api/add-session", &controllers.ApiController{}, "POST:AddSession")
	beego.Router("/api/delete-session", &controllers.ApiController{}, "POST:DeleteSession")
	beego.Router("/api/start-session", &controllers.ApiController{}, "POST:StartSession")
	beego.Router("/api/stop-session", &controllers.ApiController{}, "POST:StopSession")

	beego.Router("/api/add-asset-tunnel", &controllers.ApiController{}, "POST:AddAssetTunnel")
	beego.Router("/api/get-asset-tunnel", &controllers.ApiController{}, "GET:GetAssetTunnel")
}
