package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// Create a new cron scheduler
	c := cron.New()

	// Add a cron job to the scheduler
	_, err := c.AddFunc("@every 1s", func() {
		fmt.Println("Cron job executed at", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// Start the scheduler
	c.Start()

	// Run the scheduler for a certain duration
	time.Sleep(20 * time.Second)

	// Stop the scheduler (and jobs) gracefully
	// c.Stop()
	for {
		time.Sleep(time.Second)
	}
}
