package initx

import (
	"math"
	"time"

	"github.com/cloudslit/deca/pkg/memorycacher"
)

func InitOcspCache() *memorycacher.Cache {
	return memorycacher.New(60*time.Second, memorycacher.NoExpiration, math.MaxInt64)
}
