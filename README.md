coordtransform
[![Go Report Card](https://goreportcard.com/badge/github.com/qichengzx/coordtransform)](https://goreportcard.com/report/github.com/qichengzx/coordtransform)
----

提供百度坐标系(bd-09)、火星坐标系(国测局坐标系、gcj02)、WGS84坐标系的相互转换，基于 Go 语言，无特殊依赖。

Features
------

- [火星坐标系->百度坐标系](https://github.com/qichengzx/coordtransform/blob/ecfa225e65d96871733941ddc35ff8194c13e9f3/main.go#L32)
- [百度坐标系->WGS84坐标系](https://github.com/qichengzx/coordtransform/blob/ecfa225e65d96871733941ddc35ff8194c13e9f3/main.go#L18)
- [WGS84坐标系->火星坐标系](https://github.com/qichengzx/coordtransform/blob/ecfa225e65d96871733941ddc35ff8194c13e9f3/main.go#L43)
- [火星坐标系->WGS84坐标系](https://github.com/qichengzx/coordtransform/blob/ecfa225e65d96871733941ddc35ff8194c13e9f3/main.go#L54)
- [WGS84坐标系->百度坐标系](https://github.com/qichengzx/coordtransform/blob/ecfa225e65d96871733941ddc35ff8194c13e9f3/main.go#L71)
- [百度坐标系->WGS84坐标系](https://github.com/qichengzx/coordtransform/blob/ecfa225e65d96871733941ddc35ff8194c13e9f3/main.go#L65)

Installation
------

```bash
go get github.com/qichengzx/coordtransform
```

Quick Start
------

```Go
package main

import(
	"fmt"
	"github.com/qichengzx/coordtransform"
)

func main() {
	fmt.Println(coordTransform.BD09toGCJ02(116.404, 39.915))
	fmt.Println(coordTransform.GCJ02toBD09(116.404, 39.915))
	fmt.Println(coordTransform.WGS84toGCJ02(116.404, 39.915))
	fmt.Println(coordTransform.GCJ02toWGS84(116.404, 39.915))
	fmt.Println(coordTransform.BD09toWGS84(116.404, 39.915))
	fmt.Println(coordTransform.WGS84toBD09(116.404, 39.915))
}

```

Benchmarks
------

###### Run on MacBook Pro (13-inch, Early 2015) Go version go1.10 darwin/amd64

```Go
goos: darwin
goarch: amd64
pkg: coordTransform
BenchmarkBD09toGCJ02-4    	20000000	        84.0 ns/op
BenchmarkGCJ02toBD09-4    	20000000	        91.8 ns/op
BenchmarkWGS84toBD09-4    	 5000000	       398 ns/op
BenchmarkGCJ02toWGS84-4   	10000000	       252 ns/op
BenchmarkBD09toWGS84-4    	 5000000	       352 ns/op
BenchmarkWGS84toGCJ02-4   	 5000000	       296 ns/op
```

License
------

This project is under the MIT License.
