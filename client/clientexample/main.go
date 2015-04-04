// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/google/cadvisor/client"
	info "github.com/google/cadvisor/info/v1"
)

func staticClientExample() {
	staticClient, err := client.NewClient("http://localhost:8080/")
	if err != nil {
		glog.Errorf("tried to make client and got error %v", err)
		return
	}
	einfo, err := staticClient.EventStaticInfo("?oom_events=true&historical=true")
	if err != nil {
		glog.Errorf("got error retrieving event info: %v", err)
		return
	}
	for idx, ev := range einfo {
		glog.Infof("static einfo %v: %v", idx, ev)
	}
}

func streamingClientExample() {
	streamingClient, err := client.NewClient("http://localhost:8080/")
	if err != nil {
		glog.Errorf("tried to make client and got error %v", err)
		return
	}
	einfo := make(chan *info.Event)
	go func() {
		err = streamingClient.EventStreamingInfo("?oom_events=true", einfo)
		if err != nil {
			glog.Errorf("got error retrieving event info: %v", err)
			return
		}
	}()
	for ev := range einfo {
		glog.Infof("streaming einfo: %v", ev)
	}
}

// demonstrates how to use event clients
func main() {
	flag.Parse()
	staticClientExample()
	streamingClientExample()
}
