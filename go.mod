module github.com/beslow/go-gin-example

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/ugorji/go v1.2.1 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/sys v0.0.0-20201207223542-d4d67f95c62d // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/beslow/go-gin-example/conf => ./pkg/conf
	github.com/beslow/go-gin-example/middleware => ./middleware
	github.com/beslow/go-gin-example/models => ./models
	github.com/beslow/go-gin-example/pkg/setting => ./pkg/setting
	github.com/beslow/go-gin-example/routers => ./routers
)
