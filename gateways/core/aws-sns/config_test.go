/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws_sns

import (
	"github.com/argoproj/argo-events/pkg/apis/eventsources/v1alpha1"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var es = `
hook:
 endpoint: "/test"
 port: "8080"
 url: "myurl/test"
topicArn: "test-arn"
region: "us-east-1"
accessKey:
    key: accesskey
    name: sns
secretKey:
    key: secretkey
    name: sns
`

var esWithoutCreds = `
hook:
 endpoint: "/test"
 port: "8080"
 url: "myurl/test"
topicArn: "test-arn"
region: "us-east-1"
`

func TestParseConfig(t *testing.T) {
	convey.Convey("Given a aws-sns event source, parse it", t, func() {
		ps, err := parseEventSource(es)
		convey.So(err, convey.ShouldBeNil)
		convey.So(ps, convey.ShouldNotBeNil)
		_, ok := ps.(*v1alpha1.SNSEventSource)
		convey.So(ok, convey.ShouldEqual, true)
	})

	convey.Convey("Given a aws-sns event source without credentials, parse it", t, func() {
		ps, err := parseEventSource(esWithoutCreds)
		convey.So(err, convey.ShouldBeNil)
		convey.So(ps, convey.ShouldNotBeNil)
		_, ok := ps.(*v1alpha1.SNSEventSource)
		convey.So(ok, convey.ShouldEqual, true)
	})
}
