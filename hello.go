package hello

import (
	"rsc.io/quote/v3"
	"github.com/zhengjilei/hello/v2"
	"github.com/zhengjilei/hello/v3"
)
func Hello() string {
    return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}

func CallV2() string{
	return hello.Hello() // 注意这里是 hello: 即v2目录下的 package 名
}

func CallV3() string{
	return helloV3.Hello()
}