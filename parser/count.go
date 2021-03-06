// Package parser handles parsing log files based on the SQL execution type.
package parser

import (
	"sync"

	"github.com/busbud/tidalwave/logger"
	"github.com/busbud/tidalwave/sqlquery"
)

func countParse(query *sqlquery.QueryParams, resultsChan chan<- int, logPath string, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 0
	err := readLines(logPath, func(line *[]byte) {
		if query.ProcessLine(line) {
			count++
		}
	})

	if err != nil {
		logger.Log.Fatal(err)
	}

	resultsChan <- count
}

// Count executes a COUNT() query over log results.
// SELECT COUNT(*) FROM testapp WHERE date > '2016-10-05'
func (tp *TidalwaveParser) Count() int {
	logsLen := len(tp.LogPaths)
	resultsChan := make(chan int, logsLen)

	var wg sync.WaitGroup
	wg.Add(logsLen + 1)

	counts := []int{}
	coreLimit := make(chan bool, tp.MaxParallelism)
	go func() {
		for res := range resultsChan {
			counts = append(counts, res)
			<-coreLimit
			if len(counts) == logsLen {
				wg.Done()
			}
		}
	}()

	for i := 0; i < logsLen; i++ {
		go countParse(tp.Query, resultsChan, tp.LogPaths[i], &wg)
		coreLimit <- true
	}

	wg.Wait()

	sum := 0
	for _, count := range counts {
		sum += count
	}

	return sum
}
