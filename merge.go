package main

import (
	"github.com/efectn-forks/gobenchdata/bench"
)

func merge(files ...string) {
	histories := load(files...)
	results := make([]bench.Run, 0)
	for _, runs := range histories {
		results = append(results, runs...)
	}
	output(results)
}
