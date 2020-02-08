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

type (
	IsObject    = gqlengine.IsGraphQLObject
	IsArguments = gqlengine.IsGraphQLArguments
	IsInterface = gqlengine.IsGraphQLInterface
	IsInput     = gqlengine.IsGraphQLInput
)

type ID uint64

func (id ID) GraphQLID() {}

type Episode int

const (
	NEWHOPE Episode = iota + 1
	EMPIRE
	JEDI
)

func (e Episode) GraphQLEnumDescription() string {
	return "The episodes in the Star Wars trilogy"
}

func (e Episode) GraphQLEnumValues() gqlengine.EnumValueMapping {
	return gqlengine.EnumValueMapping{
		"NEWHOPE": {NEWHOPE, "Star Wars Episode IV: A New Hope, released in 1977."},
		"EMPIRE":  {EMPIRE, "Star Wars Episode V: The Empire Strikes Back, released in 1980."},
		"JEDI":    {JEDI, "Star Wars Episode VI: Return of the Jedi, released in 1983."},
	}
}

type Character interface {
	IsCharacter()
}

type CharacterPrototype struct {
	IsInterface `gqlDesc:"A character from the Star Wars universe"`

	ID                ID                 `json:"id" gqlDesc:"The ID of the character" gqlRequired:"true"`
	Name              string             `json:"name" gqlDesc:"The name of the character" gqlRequired:"true"`
	Friends           []Character        `json:"friends" gqlDesc:"The friends of the character, or an empty list if they have none"`
	FriendsConnection *FriendsConnection `json:"friendsConnection" gqlDesc:"The friends of the character exposed as a connection with edges" gqlRequired:"true"`
	AppearsIn         []Episode          `json:"appearsIn" gqlDesc:"The movies this character appears in" gqlRequired:"true"`
}

type FriendConnectionParams struct {
	IsArguments

	First int `json:"first"`
	After ID  `json:"after"`
}

func (p *CharacterPrototype) ResolveFriendsConnection(params *FriendConnectionParams) *FriendsConnection {
	return nil
}

type LengthUnit int

const (
	METER LengthUnit = iota + 1
	FOOT
)

func (lu LengthUnit) GraphQLEnumDescription() string { return "Units of height" }

func (lu LengthUnit) GraphQLEnumValues() gqlengine.EnumValueMapping {
	return gqlengine.EnumValueMapping{
		"METER": {METER, "The standard unit around the world"},
		"FOOT":  {FOOT, "Primarily used in the United States"},
	}
}

type Human struct {
	IsObject `gqlDesc:"A humanoid creature from the Star Wars universe"`

	ID                ID                 `json:"id" gqlDesc:"The ID of the human" gqlRequired:"true"`
	Name              string             `json:"name" gqlDesc:"What this human calls themselves" gqlRequired:"true"`
	HomePlanet        string             `json:"homePlanet" gqlDesc:"The home planet of the human, or null if unknown"`
	Height            float32            `json:"height" gqlDesc:"Height in the preferred unit, default is meters"`
	Mass              float32            `json:"mass" gqlDesc:"Mass in kilograms, or null if unknown"`
	Friends           []Character        `json:"friends" gqlDesc:"This human's friends, or an empty list if they have none"`
	FriendsConnection *FriendsConnection `json:"friendsConnection" gqlDesc:"The friends of the human exposed as a connection with edges" gqlRequired:"true"`
	AppearsIn         []Episode          `json:"appearsIn" gqlDesc:"The movies this human appears in" gqlRequired:"true"`
	Starships         []*Starship        `json:"starships" gqlDesc:"A list of starships this person has piloted, or an empty list if none"`

	// internal fields
	friendIds []ID
	starships []ID
}

func (h *Human) IsCharacter() {}

type HumanHeightParams struct {
	gqlengine.IsGraphQLArguments
	Unit LengthUnit `json:"unit" gqlDesc:"height unit" gqlDefault:"1"`
}

func (h *Human) ResolveHeight(params *HumanHeightParams) float32 {
	if h.Height < 0 {
		return 0
	}
	if params.Unit == METER {
		return h.Height
	}
	return h.Height * 3.2808399
}

func (h *Human) ResolveMass() float32 {
	if h.Mass < 0 {
		return 0
	}
	return h.Mass
}

func (h *Human) ResolveFriends() []Character {
	return resolveFriends(h.friendIds)
}

func (h *Human) ResolveFriendsConnection(params *FriendConnectionParams) *FriendsConnection {
	return resolveFriendConnection(params, h.friendIds)
}

func (h *Human) ResolveStarships() []*Starship {
	ships := make([]*Starship, len(h.starships))
	for i, id := range h.starships {
		ships[i] = starshipMapping[id]
	}
	return ships
}

type Droid struct {
	IsObject          `gqlDesc:"An autonomous mechanical character in the Star Wars universe"`
	ID                ID                 `json:"id" gqlDesc:"The ID of the droid" gqlRequired:"true"`
	Name              string             `json:"name" gqlDesc:"What others call this droid" gqlRequired:"true"`
	Friends           []Character        `json:"friends" gqlDesc:"This droid's friends, or an empty list if they have none"`
	FriendsConnection *FriendsConnection `json:"friendsConnection" gqlDesc:"The friends of the droid exposed as a connection with edges" gqlRequired:"true"`
	AppearsIn         []Episode          `json:"appearsIn" gqlDesc:"The movies this droid appears in" gqlRequired:"true"`
	PrimaryFunction   string             `json:"primaryFunction" gqlDesc:"This droid's primary function"`

	friendIds []ID
}

func (d *Droid) IsCharacter() {}

func (d *Droid) ResolveFriends() []Character {
	return resolveFriends(d.friendIds)
}

func (d *Droid) ResolveFriendsConnection(params *FriendConnectionParams) *FriendsConnection {
	return resolveFriendConnection(params, d.friendIds)
}

type FriendsConnection struct {
	IsObject `gqlDesc:"A connection object for a character's friends"`

	TotalCount int           `json:"totalCount" gqlDesc:"The total number of friends"`
	Edges      []*FriendEdge `json:"edges" gqlDesc:"The edges for each of the character's friends."`
	Friends    []Character   `json:"friends" gqlDesc:"A list of the friends, as a convenience when edges are not needed."`
	PageInfo   *PageInfo     `json:"pageInfo" gqlDesc:"Information for paginating this connection" gqlRequired:"true"`
}

type FriendEdge struct {
	IsObject `gqlDesc:"An edge object for a character's friends"`

	Cursor ID        `json:"cursor" gqlDesc:"A cursor used for pagination" gqlRequired:"true"`
	Node   Character `json:"node" gqlDesc:"The character represented by this friendship edge"`
}

type PageInfo struct {
	IsObject `gqlDesc:"Information for paginating this connection"`

	StartCursor *ID  `json:"startCursor"`
	EndCursor   *ID  `json:"endCursor"`
	HasNextPage bool `json:"hasNextPage" gqlRequired:"true"`
}

type Review struct {
	IsObject `gqlDesc:"Represents a review for a movie"`

	Episode    Episode `json:"episode" gqlDesc:"The movie"`
	Stars      int     `json:"stars" gqlDesc:"The number of stars this review gave, 1-5" gqlRequired:"true"`
	Commentary string  `json:"commentary" gqlDesc:"Comment about the movie"`
}

type ReviewInput struct {
	IsInput `gqlDesc:"The input object sent when someone is creating a new review"`

	Stars         int         `json:"stars" gqlDesc:"0-5 stars" gqlRequired:"true"`
	Commentary    string      `json:"commentary" gqlDesc:"Comment about the movie, optional"`
	FavoriteColor *ColorInput `json:"favorite_color" gqlDesc:"Favorite color, optional"`
}

type ColorInput struct {
	IsInput `gqlDesc:"The input object sent when passing in a color"`

	Red   int `json:"red" gqlRequired:"true"`
	Green int `json:"green" gqlRequired:"true"`
	Blue  int `json:"blue" gqlRequired:"true"`
}

type Starship struct {
	IsObject

	ID          ID          `json:"id" gqlDesc:"The ID of the starship" gqlRequired:"true"`
	Name        string      `json:"name" gqlDesc:"The name of the starship" gqlRequired:"true"`
	Length      float32     `json:"length" gqlDesc:"Length of the starship, along the longest axis"`
	Coordinates [][]float32 `json:"coordinates" gqlDesc:"coordinates" gqlRequired:"true"`
}

type SearchResult interface {
	IsSearchResult()
}

func (h *Human) IsSearchResult()    {}
func (d *Droid) IsSearchResult()    {}
func (s *Starship) IsSearchResult() {}
