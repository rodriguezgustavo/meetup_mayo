package cache4


import (
	"testing"
	"github.com/mercadolibre/meetup/cache/cachetest"
)

var httpGetBody = cachetest.HTTPGetBody

/*func Test(t *testing.T) {
	m := New(httpGetBody)
	cachetest.Sequential(t, m)
}*/

// NOTE: not concurrency-safe!  Test fails.
func TestConcurrent(t *testing.T) {
	m := New(httpGetBody)
	cachetest.Concurrent(t, m)
}