package main

import (
	"errors"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	logger.Info("Starting the Task Processor")

	tasks := []string{"task1", "task2", "task3", "task4", "task5"}

	for {
		for _, task := range tasks {
			logger.Info("Processing task", zap.String("task", task))

			err := processTask(task)
			if err != nil {
				logger.Error("Task processing failed",
					zap.String("task", task),
					zap.Error(err))
			} else {
				logger.Info("Task completed successfully", zap.String("task", task))
			}
		}
	}
}

// processTask simulates processing a task and randomly returns an error
func processTask(task string) error {
	// Simulate random processing time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1500)))

	// Simulate a random error for demo purposes
	if rand.Intn(10) > 7 { // 30% chance of failure
		return errors.New("simulated error during task processing")
	}
	return nil
}
