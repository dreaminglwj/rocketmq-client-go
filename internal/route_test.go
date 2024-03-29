/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.dreaminglwj.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package internal

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"

	"github.com/dreaminglwj/rocketmq-client-go/internal/remote"
	"github.com/dreaminglwj/rocketmq-client-go/primitive"
)

func TestQueryTopicRouteInfoFromServer(t *testing.T) {
	Convey("marshal of TraceContext", t, func() {

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		remotingCli := remote.NewMockRemotingClient(ctrl)

		addr, err := primitive.NewNamesrvAddr("1.1.1.1:8880", "1.1.1.2:8880", "1.1.1.3:8880")
		assert.Nil(t, err)

		namesrv, err := NewNamesrv(addr)
		assert.Nil(t, err)
		namesrv.nameSrvClient = remotingCli

		Convey("When marshal producer trace data", func() {

			count := 0
			remotingCli.EXPECT().InvokeSync(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
				func(ctx context.Context, addr string, request *remote.RemotingCommand, timeout time.Duration) (*remote.RemotingCommand, error) {
					count++
					if count < 3 {
						return nil, errors.New("not existed")
					}
					return &remote.RemotingCommand{
						Code: ResTopicNotExist,
					}, nil
				}).Times(3)

			data, err := namesrv.queryTopicRouteInfoFromServer("notexisted")
			assert.Nil(t, data)
			assert.Equal(t, ErrTopicNotExist, err)
		})
	})
}
