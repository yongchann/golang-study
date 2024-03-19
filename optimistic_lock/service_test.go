package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"sync"
	"testing"
	"time"
)

func TestLockService_updateWithOptimisticLock(t *testing.T) {
	db := initDB()

	tests := []struct {
		name          string
		productId     uint64
		iterations    int
		sleepInterval time.Duration
	}{
		{
			name:          "1ms interval",
			productId:     1,
			iterations:    100,
			sleepInterval: time.Millisecond * 1,
		},
		{
			name:          "2ms interval",
			productId:     2,
			iterations:    100,
			sleepInterval: time.Millisecond * 2,
		},
		{
			name:          "4ms interval",
			productId:     3,
			iterations:    100,
			sleepInterval: time.Millisecond * 4,
		},
		{
			name:          "8ms interval",
			productId:     4,
			iterations:    100,
			sleepInterval: time.Millisecond * 8,
		},
		{
			name:          "16ms interval",
			productId:     5,
			iterations:    100,
			sleepInterval: time.Millisecond * 16,
		},
		{
			name:          "32ms interval",
			productId:     6,
			iterations:    100,
			sleepInterval: time.Millisecond * 32,
		},
		{
			name:          "64ms interval",
			productId:     7,
			iterations:    100,
			sleepInterval: time.Millisecond * 64,
		},
		{
			name:          "128ms interval",
			productId:     8,
			iterations:    100,
			sleepInterval: time.Millisecond * 128,
		},
	}

	type TestResult struct {
		Name           string
		FailPercentage float64
	}

	var testResults []TestResult

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := LockService{
				lockRepository: NewLockRepositoryImpl(db),
			}

			failCnt := 0
			wg := sync.WaitGroup{}
			wg.Add(tt.iterations)

			for i := 0; i < tt.iterations; i++ {
				time.Sleep(tt.sleepInterval)
				go func() {
					start := time.Now()
					if err := svc.updateWithOptimisticLock(tt.productId); err != nil {
						failCnt += 1
					}
					end := time.Since(start)
					fmt.Println(end.String())
					wg.Done()
				}()
			}
			wg.Wait()
			testResults = append(testResults, TestResult{tt.name, float64(failCnt*100) / float64(tt.iterations)})
			t.Logf("failCnt: %3d, failfail percentage: %.01f%%\n", failCnt, float64(failCnt*100)/float64(tt.iterations))
		})
	}

	// Create a bar chart
	p := plot.New()
	p.Title.Text = "Test Results"
	p.Y.Label.Text = "Fail Percentage"
	p.X.Label.Text = "Interval"

	// Create a bar plotter
	bars := make(plotter.Values, len(testResults))
	tickLabels := make([]string, len(testResults))
	for i, result := range testResults {
		bars[i] = result.FailPercentage
		tickLabels[i] = result.Name
	}

	bp, err := plotter.NewBarChart(bars, vg.Points(50))
	if err != nil {
		log.Fatal(err)
	}
	bp.LineStyle.Width = vg.Length(0)
	bp.Color = color.RGBA{R: 255, B: 128, A: 255}

	// Add the bar plotter to the plot
	p.Add(bp)

	// Set the X axis tick labels
	p.NominalX(tickLabels...)

	// Save the plot to a file
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "test_results.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bar chart saved as test_results.png")
}
