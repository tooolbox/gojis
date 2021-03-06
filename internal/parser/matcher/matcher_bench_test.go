package matcher

import (
	"testing"
)

var result interface{}

func BenchmarkString(b *testing.B) {
	matcher := String("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var r bool

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = matcher.Matches('2') &&
			matcher.Matches('d') &&
			matcher.Matches('u') &&
			matcher.Matches('F') &&
			matcher.Matches('U')
	}

	result = r
}

func BenchmarkMerge(b *testing.B) {
	matcher := Merge(
		String("0123456789"),
		String("abcdefghijklmnopqrstuvwxyz"),
		String("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
	)
	var r bool

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r = matcher.Matches('2') &&
			matcher.Matches('d') &&
			matcher.Matches('u') &&
			matcher.Matches('F') &&
			matcher.Matches('U')
	}

	result = r
}
