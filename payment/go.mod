module github.com/pwera/grpc-micros-payment

go 1.20

require (
	github.com/golang/protobuf v1.5.3
	github.com/mwitkow/go-proto-validators v0.3.2
<<<<<<< HEAD
	github.com/pwera/grpc-micros-commons v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.2.3
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.46.1
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4
=======
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)

require (
<<<<<<< HEAD
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.2.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.opentelemetry.io/otel/sdk v1.21.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231002182017-d307bd883b97 // indirect
)

replace github.com/pwera/grpc-micros-commons => ../commons
=======
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
>>>>>>> d5d6d859f89eecf70f457a12f02ef3d2b3daf9e4
