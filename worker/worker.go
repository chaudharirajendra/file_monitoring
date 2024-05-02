package worker

import (
	"file_monitoring/database"
	"file_monitoring/file"
	"log"
	"sync"
)

func StartWorkers(concurrency int, wg *sync.WaitGroup, fileCh <-chan string, db *database.SQLiteDB) {

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for filePath := range fileCh {
				size, err := file.GetFileSize(filePath)
				if err != nil {
					log.Printf("Error reading file %s: %v\n", filePath, err)
				} else {
					log.Printf("Total number of bytes in file %s: %d\n", filePath, size)
					db.InsertFileInfo(filePath, size)
				}
			}
		}()
	}
}
