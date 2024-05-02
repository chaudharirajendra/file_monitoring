package main

import (
	"file_monitoring/config"
	"file_monitoring/database"
	"file_monitoring/file"
	"file_monitoring/worker"
	"log"
	"sync"
)

func main() {

	// Create an instance of GodotEnvLoader
	loader := &config.GodotEnvLoader{}

	// Create a new config using the loader
	conf, err := config.NewConfig(loader)
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize database
	db, err := database.NewSQLiteDB(conf.StoragePath)
	if err != nil {
		log.Fatal("Error loading the db obj:", err)
	}
	defer db.Close()

	// Initialize file watcher
	watcher, err := file.NewWatcher(conf.TargetDir)
	if err != nil {
		log.Fatal("Error loading watcher obj:", err)
	}

	// Start workers
	wg := &sync.WaitGroup{}
	worker.StartWorkers(conf.Concurrency, wg, watcher.FileChan(), db)

	// Start watching files
	watcher.Start()

	// Wait for all workers to finish
	wg.Wait()
}
