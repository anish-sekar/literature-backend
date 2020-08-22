package utils

import (
	"math/rand"
  "time"
  "github.com/anish-sekar/literature-backend/models"
)

func ShuffleCards(vals []string) []string {
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

func ShufflePlayers(vals []*models.Player) []*models.Player{
  r := rand.New(rand.NewSource(time.Now().Unix()))
    ret := make([]*models.Player, len(vals))
    n := len(vals)
    for i := 0; i < n; i++ {
      randIndex := r.Intn(len(vals))
      ret[i] = vals[randIndex]
      vals = append(vals[:randIndex], vals[randIndex+1:]...)
    }
    return ret


}