// Copyright 2020 凯斐德科技（杭州）有限公司 (Karfield Technology, ltd.)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gqlengine/gqlengine"
	"github.com/gqlengine/playground"
)

func main() {
	engine := gqlengine.NewEngine(gqlengine.Options{
		Tracing: true, // enable tracing extensions
	})

	if err := initTypes(engine); err != nil {
		panic(err)
	}

	if err := initResolvers(engine); err != nil {
		panic(err)
	}

	if err := engine.Init(); err != nil {
		panic(err)
	}

	playground.SetEndpoints("/api/graphql", "/api/graphql/subscriptions")

	r := mux.NewRouter()
	r.HandleFunc("/api/graphql", engine.ServeHTTP)
	r.HandleFunc("/api/graphql/subscriptions", engine.ServeWebsocket)
	r.PathPrefix("/api/graphql/playground").Handler(playground.GetHandle("/api/graphql/playground"))

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(AssetFile())))

	println("open demo page at: http://localhost:9996/")
	println("open playground http://localhost:9996/api/graphql/playground/")
	if err := http.ListenAndServe(":9996", r); err != nil {
		panic(err)
	}
}
