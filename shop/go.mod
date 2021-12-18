module shop

go 1.12

require (
	common v0.0.0
	github.com/agiledragon/gomonkey v2.0.2+incompatible
	github.com/go-sql-driver/mysql v1.6.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/smartystreets/goconvey v1.7.2
	github.com/tal-tech/go-zero v1.2.4
	gopkg.in/yaml.v2 v2.4.0
)

replace common => ../common
