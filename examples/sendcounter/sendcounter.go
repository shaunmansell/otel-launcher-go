// Copyright Lightstep Authors
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

// For example:
//
//   OTEL_EXPORTER_OTLP_METRIC_ENDPOINT=localhost:17000
//   OTEL_EXPORTER_OTLP_METRIC_INSECURE=true \
//   LS_SERVICE_NAME=foo \
//   LS_ACCESS_TOKEN=12341234123412341234123412341234 \
//   go run sendcounter.go

package main

import (
	"context"

	"github.com/lightstep/otel-launcher-go/launcher"
	"go.opentelemetry.io/otel/metric"
	metricglobal "go.opentelemetry.io/otel/metric/global"
)

func main() {
	defer launcher.ConfigureOpentelemetry().Shutdown()

	ctx := context.Background()
	meter := metric.Must(metricglobal.Meter("commandline"))

	c1 := meter.NewInt64Counter("counter")
	c1.Add(ctx, 100)
}
