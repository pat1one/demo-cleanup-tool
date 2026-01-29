package main

var version = "dev"

import (
var version = "dev"
"flag"
"fmt"
"log"
"os"
"path/filepath"
"time"
)

type CleanupStats struct {
TotalFiles   int
DeletedFiles int
FreedSpace   int64
}

func main() {
demoPath := flag.String("path", "./tmp_demos", "Path to demo files directory")
daysOld := flag.Int("days", 7, "Delete files older than N days")
dryRun := flag.Bool("dry-run", false, "Show what would be deleted without deleting")
showVersion := flag.Bool("version", false, "Show version")
	showVersion := flag.Bool("version", false, "Show version")
flag.Parse()
if *showVersion {
fmt.Println(version)
return
}

	if *showVersion {
		fmt.Println(version)
		return
	}

fmt.Printf("Scanning: %s\n", *demoPath)
fmt.Printf("Deleting files older than %d days\n", *daysOld)

if *dryRun {
fmt.Println("DRY RUN MODE - no files will be deleted")
}

stats := CleanupStats{}
cutoffTime := time.Now().AddDate(0, 0, -*daysOld)

err := filepath.Walk(*demoPath, func(path string, info os.FileInfo, err error) error {
if err != nil {
return err
}

if info.IsDir() {
return nil
}

stats.TotalFiles++

if filepath.Ext(path) == ".dem" && info.ModTime().Before(cutoffTime) {
size := info.Size()

if *dryRun {
fmt.Printf("Would delete: %s (%.2f MB)\n", path, float64(size)/1024/1024)
} else {
if err := os.Remove(path); err != nil {
log.Printf("Error deleting %s: %v\n", path, err)
return nil
}
fmt.Printf("Deleted: %s (%.2f MB)\n", path, float64(size)/1024/1024)
}

stats.DeletedFiles++
stats.FreedSpace += size
}

return nil
})

if err != nil {
log.Fatal(err)
}

fmt.Println("\nSummary:")
fmt.Printf("Total files scanned: %d\n", stats.TotalFiles)
fmt.Printf("Files deleted: %d\n", stats.DeletedFiles)
fmt.Printf("Space freed: %.2f MB\n", float64(stats.FreedSpace)/1024/1024)
}
