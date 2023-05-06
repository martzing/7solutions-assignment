package service

import (
	"testing"

	"github.com/martzing/7solutions-assignment/3-pie-fire-dire/configs"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetBeefSummary(t *testing.T) {
	t.Run("Test summary complete", func(t *testing.T) {
		beefUrl := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
		configs.BeefUrl = &beefUrl
		got, err := GetBeefSummary()
		assert.NotEqual(t, nil, got)
		assert.Equal(t, nil, err)
	})

	t.Run("Test summary error", func(t *testing.T) {
		beefUrl := "https://url.com"
		configs.BeefUrl = &beefUrl
		got, err := GetBeefSummary()
		assert.NotEqual(t, nil, err)
		assert.Equal(t, nil, got)
	})

}
