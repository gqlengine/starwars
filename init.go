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
	"github.com/gqlengine/gqlengine"
)

func initTypes(engine *gqlengine.Engine) error {
	if err := engine.PreRegisterInterface((*Character)(nil), CharacterPrototype{}); err != nil {
		return err
	}
	if _, err := engine.RegisterObject(Human{}); err != nil {
		return err
	}
	if _, err := engine.RegisterObject(Droid{}); err != nil {
		return err
	}
	if err := engine.PreRegisterUnion((*SearchResult)(nil), Human{}, Droid{}, Starship{}); err != nil {
		return err
	}

	return nil
}

func initResolvers(engine *gqlengine.Engine) error {
	engine.NewQuery(GetHero).
		Name("hero").
		Description("Allows us to fetch the undisputed hero of the Star Wars trilogy, R2-D2.")

	engine.NewQuery(Reviews).
		Description("Allows us to fetch the ephemeral reviews for each episode")

	engine.NewMutation(CreateReviews).
		Description("set reviews")

	engine.NewQuery(Search)

	engine.NewQuery(GetHuman).
		Name("human")

	engine.NewQuery(GetDroid).
		Name("droid")

	engine.NewQuery(GetStartship).
		Name("starship")

	return nil
}
