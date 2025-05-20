package main

import (
	"errors"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

func main() {
	// Initialize Zap logger
	logger, err := zap.NewProduction() // Creates a production-ready logger
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync() // Flushes buffer, if any, before program exit

	logger.Info("Starting the Task Processor")

	// Simulate a list of tasks
	tasks := []string{"task1", "task2", "task3", "task4", "task5"}

	// Process each task
	for _, task := range tasks {
		logger.Info("Processing task", zap.String("task", task))

		// Simulate task processing
		err := processTask(task)
		if err != nil {
			logger.Error("Task processing failed",
				zap.String("task", task),
				zap.Error(err))
		} else {
			logger.Info("Task completed successfully", zap.String("task", task))
		}
	}

	logger.Info("Finished processing all tasks")
}

// processTask simulates processing a task and randomly returns an error
func processTask(task string) error {
	// Simulate random processing time
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))

	// Simulate a random error for demo purposes
	if rand.Intn(10) > 7 { // 30% chance of failure
		return errors.New("simulated error during task processing")
	}
	return nil
}
