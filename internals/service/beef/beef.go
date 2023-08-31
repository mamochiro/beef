package beef

import (
	"context"
	"regexp"
	"strings"
	"sync"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func (s Service) BeefSummary(ctx context.Context) (map[string]int32, error) {
	words, err := s.restRepo.BeefSummary(ctx)
	if err != nil {
		return nil, err
	}
	res := map[string]int32{}
	// for _, w := range *words {
	// 	str := strings.ToLower(clearString(w))
	// 	_, ok := res[str]
	// 	if ok {
	// 		res[str] += 1
	// 	} else {
	// 		res[str] = 1
	// 	}
	// }
	var mu sync.Mutex
	lop.ForEach(lo.FromPtr(words), func(w string, _ int) {
		mu.Lock()
		defer mu.Unlock()
		str := strings.ToLower(clearString(w))
		_, ok := res[str]
		if ok {
			res[str] += 1
		} else {
			res[str] = 1
		}
	})
	return res, nil
}
