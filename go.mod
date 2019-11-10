module github.com/zhengjilei/hello

go 1.12

require (
	github.com/zhengjilei/hello/v2 v2.0.0-00010101000000-000000000000
	github.com/zhengjilei/hello/v3 v3.0.0-00010101000000-000000000000
	rsc.io/quote/v3 v3.1.0
)

replace github.com/zhengjilei/hello/v2 => ./v2

replace github.com/zhengjilei/hello/v3 => ./v3
