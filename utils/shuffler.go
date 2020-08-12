package utils

import (
	"math/rand"
	"time"

)

func Shuffle(vals []string) []string {
    r := rand.New(rand.NewSource(time.Now().Unix()))
    ret := make([]string, len(vals))
    n := len(vals)
    for i := 0; i < n; i++ {
      randIndex := r.Intn(len(vals))
      ret[i] = vals[randIndex]
      vals = append(vals[:randIndex], vals[randIndex+1:]...)
    }
    return ret
  }