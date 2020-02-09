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

var humans = []*Human{
	{
		ID:         1000,
		Name:       "Luke Skywalker",
		friendIds:  []ID{1002, 1003, 2000, 2001},
		AppearsIn:  []Episode{NEWHOPE, EMPIRE, JEDI},
		HomePlanet: "Tatooine",
		Height:     1.72,
		Mass:       77,
		starships:  []ID{3001, 3003},
	},
	{
		ID:        1001,
		Name:      "Han Solo",
		friendIds: []ID{1000, 1003, 2001},
		AppearsIn: []Episode{NEWHOPE, EMPIRE, JEDI},
		Height:    1.8,
		Mass:      80,
		starships: []ID{3000, 3003},
	},
	{
		ID:         1003,
		Name:       "Leia Organa",
		friendIds:  []ID{1000, 1002, 2000, 2001},
		AppearsIn:  []Episode{NEWHOPE, EMPIRE, JEDI},
		HomePlanet: "Alderaan",
		Height:     1.5,
		Mass:       49,
	},
	{
		ID:        1004,
		Name:      "Wilhuff Tarkin",
		friendIds: []ID{1001},
		AppearsIn: []Episode{NEWHOPE},
		Height:    1.8,
		Mass:      -1,
	},
}

var droids = []*Droid{
	{
		ID:              2000,
		Name:            "C-3PO",
		friendIds:       []ID{1000, 1002, 1003, 2001},
		AppearsIn:       []Episode{NEWHOPE, EMPIRE, JEDI},
		PrimaryFunction: "Protocol",
	},
	{
		ID:              2001,
		Name:            "R2-D2",
		friendIds:       []ID{1000, 1002, 1003},
		AppearsIn:       []Episode{NEWHOPE, EMPIRE, JEDI},
		PrimaryFunction: "Astromech",
	},
}

var characters = map[ID]Character{}

func init() {
	for _, h := range humans {
		characters[h.ID] = h
	}
	for _, h := range droids {
		characters[h.ID] = h
	}
}

var starships = []*Starship{
	{
		ID:     3000,
		Name:   "Millenium Falcon",
		Length: 34.37,
	},
	{
		ID:     3001,
		Name:   "X-Wing",
		Length: 12.5,
	},
	{
		ID:     3002,
		Name:   "TIE Advanced x1",
		Length: 9.2,
	},
	{
		ID:     3003,
		Name:   "Imperial shuttle",
		Length: 20,
	},
}

var starshipMapping = map[ID]*Starship{}

func init() {
	for _, ss := range starships {
		starshipMapping[ss.ID] = ss
	}
}

var reviews = map[Episode][]*Review{
	NEWHOPE: {
		{
			Stars:      5,
			Commentary: "my favorite episode",
		},
	},
}
