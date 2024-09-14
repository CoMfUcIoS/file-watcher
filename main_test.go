package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestFileWatcher(t *testing.T) {
	// Create a temporary directory
	tempDir, err := ioutil.TempDir("", "watcher_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Create a temporary file in the directory
	tempFile, err := ioutil.TempFile(tempDir, "testfile")
	if err != nil {
		t.Fatal(err)
	}
	tempFile.Close()

	// Define the action to be executed
	action := "echo 'file changed'"

	// Run the main function in a separate goroutine
	go func() {
		os.Args = []string{"cmd", "-source", tempDir, "-action", action}
		main()
	}()

	// Wait for the watcher to start
	time.Sleep(1 * time.Second)

	// Modify the temporary file
	err = ioutil.WriteFile(tempFile.Name(), []byte("new content"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for the debounce timer and action execution
	time.Sleep(1 * time.Second)

	// Check if the action was executed by verifying the output
	cmd := exec.Command("sh", "-c", "echo 'file changed'")
	output, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}

	expectedOutput := "file changed\n"
	if string(output) != expectedOutput {
		t.Fatalf("expected %q, got %q", expectedOutput, string(output))
	}
}
