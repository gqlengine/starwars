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
	"regexp"
)

func GetHero(params *struct {
	IsArguments
	Episode Episode `json:"episode"`
}) Character {
	if params.Episode == EMPIRE {
		return humans[0]
	}
	return droids[1]
}

func Reviews(params *struct {
	IsArguments
	Episode Episode `json:"episode" gqlRequired:"true"`
}) []*Review {
	return reviews[params.Episode]
}

func CreateReviews(params *struct {
	IsArguments
	Episode Episode      `json:"episode" gqlRequired:"true"`
	Review  *ReviewInput `json:"review" gqlRequired:"true"`
}) *Review {
	review := &Review{
		Episode:    params.Episode,
		Stars:      params.Review.Stars,
		Commentary: params.Review.Commentary,
	}
	reviews[params.Episode] = append(reviews[params.Episode], review)
	return review
}

func GetHuman(params *struct {
	IsArguments
	ID ID `json:"id" gqlRequired:"true"`
}) *Human {
	if c, ok := characters[params.ID]; ok {
		if h, ok := c.(*Human); ok {
			return h
		}
	}
	return nil
}

func GetDroid(params *struct {
	IsArguments
	ID ID `json:"id" gqlRequired:"true"`
}) *Droid {
	if c, ok := characters[params.ID]; ok {
		if h, ok := c.(*Droid); ok {
			return h
		}
	}
	return nil
}

func GetStartship(params *struct {
	IsArguments
	ID ID `json:"id" gqlRequired:"true"`
}) *Starship {
	for _, ss := range starships {
		if ss.ID == params.ID {
			return ss
		}
	}
	return nil
}

func Search(params *struct {
	IsArguments
	Text string `json:"text" gqlDesc:"input regex pattern" gqlRequired:"true"`
}) (
	results []SearchResult,
	err error,
) {
	re, e := regexp.Compile(params.Text)
	if e != nil {
		err = e
		return
	}
	for _, h := range humans {
		if re.MatchString(h.Name) || re.MatchString(h.HomePlanet) {
			results = append(results, h)
		}
	}
	for _, d := range droids {
		if re.MatchString(d.Name) || re.MatchString(d.PrimaryFunction) {
			results = append(results, d)
		}
	}
	for _, s := range starships {
		if re.MatchString(s.Name) {
			results = append(results, s)
		}
	}
	return
}

func resolveFriends(friendIds []ID) []Character {
	if len(friendIds) > 0 {
		friends := make([]Character, len(friendIds))
		for i, id := range friendIds {
			friends[i] = characters[id]
		}
		return friends
	}
	return nil
}

func resolveFriendConnection(params *FriendConnectionParams, friendIds []ID) *FriendsConnection {
	if friendIds == nil {
		friendIds = []ID{}
	}

	var ids []ID
	idx := 0
	if params.After > 0 {
		for i, id := range friendIds {
			if id == params.After {
				idx = i
				break
			}
		}
	}
	if idx+params.First < len(friendIds) {
		ids = friendIds[idx : idx+params.First]
	} else {
		ids = friendIds[idx:]
	}

	friends := make([]Character, len(ids))
	edges := make([]*FriendEdge, len(ids))
	for i, id := range ids {
		friends[i] = characters[id]
		edges[i] = &FriendEdge{
			Cursor: id,
			Node:   characters[id],
		}
	}

	var (
		startCursor *ID
		endCursor   *ID
	)
	if len(ids) > 0 {
		startCursor = &ids[0]
		endCursor = &ids[len(ids)-1]
	}

	return &FriendsConnection{
		TotalCount: len(friendIds),
		Edges:      edges,
		Friends:    friends,
		PageInfo: &PageInfo{
			StartCursor: startCursor,
			EndCursor:   endCursor,
			HasNextPage: idx+params.First < len(friendIds),
		},
	}
}
