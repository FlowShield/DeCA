package initx

import (
	"github.com/cloudslit/newca/pkg/memorycacher"
	"math"
	"time"
)

func InitOcspCache() *memorycacher.Cache {
	return memorycacher.New(60*time.Second, memorycacher.NoExpiration, math.MaxInt64)
}
