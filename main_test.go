package coordTransform

import (
	"testing"
)

func BenchmarkBD09toGCJ02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BD09toGCJ02(116.404, 39.915)
	}
}

func BenchmarkGCJ02toBD09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCJ02toBD09(116.404, 39.915)
	}
}

func BenchmarkWGS84toBD09(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WGS84toBD09(116.404, 39.915)
	}
}

func BenchmarkGCJ02toWGS84(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCJ02toWGS84(116.404, 39.915)
	}
}

func BenchmarkBD09toWGS84(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BD09toWGS84(116.404, 39.915)
	}
}

func BenchmarkWGS84toGCJ02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WGS84toGCJ02(116.404, 39.915)
	}
}
