package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Checkmarx/kics/pkg/scan"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestLoggerMultithreaded(t *testing.T) {
	// Save original logger to restore later
	originalLogger := log.Logger

	// Create a buffer to capture logs
	var buf bytes.Buffer
	testLogger := zerolog.New(&buf).With().Timestamp().Logger()
	log.Logger = testLogger

	defer func() {
		log.Logger = originalLogger
	}()

	const numThreads = 10
	const logsPerThread = 5
	
	var wg sync.WaitGroup
	threadLogs := make(map[int][]string)
	var mu sync.Mutex
	
	// Launch multiple goroutines, each with unique thread-specific data
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			
			// Create unique extra info for this thread
			extraInfos := map[string]string{
				"thread_id":    fmt.Sprintf("thread_%d", threadID),
				"thread_data": fmt.Sprintf("data_for_thread_%d", threadID),
			}
			
			// Initialize config with thread-specific data (this calls initializeConfig at line 40)
			params := scan.GetDefaultParameters("", extraInfos, false)
			if params == nil {
				t.Errorf("Failed to get default parameters for thread %d", threadID)
				return
			}
			
			// Generate logs with thread-specific content
			var threadLogLines []string
			for j := 0; j < logsPerThread; j++ {
				logMsg := fmt.Sprintf("Log message %d from thread %d", j, threadID)
				log.Info().Str("message_id", fmt.Sprintf("msg_%d_%d", threadID, j)).Msg(logMsg)
				threadLogLines = append(threadLogLines, logMsg)
			}
			
			// Store the expected logs for this thread
			mu.Lock()
			threadLogs[threadID] = threadLogLines
			mu.Unlock()
			
			// Small delay to allow logs to be written
			time.Sleep(10 * time.Millisecond)
		}(i)
	}
	
	// Wait for all threads to complete
	wg.Wait()
	
	// Give a bit more time for all logs to be written
	time.Sleep(50 * time.Millisecond)
	
	// Parse the captured logs
	logOutput := buf.String()
	logLines := strings.Split(strings.TrimSpace(logOutput), "\n")
	
	// Verify that each thread's logs contain only its own data
	threadLogCounts := make(map[int]int)
	
	for _, line := range logLines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		
		// Parse JSON log entry
		var logEntry map[string]interface{}
		if err := json.Unmarshal([]byte(line), &logEntry); err != nil {
			t.Logf("Failed to parse log line: %s", line)
			continue
		}
		
		// Check if this log contains thread-specific data
		if threadIDStr, exists := logEntry["thread_id"]; exists {
			if threadIDInterface, ok := threadIDStr.(string); ok {
				// Extract thread ID from "thread_X" format
				var threadID int
				if _, err := fmt.Sscanf(threadIDInterface, "thread_%d", &threadID); err == nil {
					threadLogCounts[threadID]++
					
					// Verify that the log contains the correct thread data
					if threadDataStr, dataExists := logEntry["thread_data"]; dataExists {
						expectedData := fmt.Sprintf("data_for_thread_%d", threadID)
						assert.Equal(t, expectedData, threadDataStr, 
							"Log should contain correct thread-specific data for thread %d", threadID)
					}
					
					// Verify that the log doesn't contain data from other threads
					for otherThreadID := 0; otherThreadID < numThreads; otherThreadID++ {
						if otherThreadID != threadID {
							wrongData := fmt.Sprintf("data_for_thread_%d", otherThreadID)
							assert.NotContains(t, line, wrongData, 
								"Thread %d log should not contain data from thread %d", threadID, otherThreadID)
						}
					}
				}
			}
		}
	}
	
	// Verify that we received logs from all threads
	assert.Equal(t, numThreads, len(threadLogCounts), "Should have received logs from all threads")
	
	// Verify that each thread produced the expected number of logs
	for threadID := 0; threadID < numThreads; threadID++ {
		count, exists := threadLogCounts[threadID]
		assert.True(t, exists, "Should have logs from thread %d", threadID)
		assert.GreaterOrEqual(t, count, logsPerThread, 
			"Thread %d should have produced at least %d logs, got %d", threadID, logsPerThread, count)
	}
	
	// Additional verification: ensure no cross-contamination
	for _, line := range logLines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		
		// Count how many different thread IDs appear in a single log line
		threadMatches := 0
		for i := 0; i < numThreads; i++ {
			if strings.Contains(line, fmt.Sprintf("thread_%d", i)) {
				threadMatches++
			}
		}
		
		// Each log line should contain data from exactly one thread (or be a system log)
		if threadMatches > 0 {
			assert.Equal(t, 1, threadMatches, 
				"Log line should not contain data from multiple threads: %s", line)
		}
	}
}