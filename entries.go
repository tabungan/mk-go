package mk

import "sync"

type Entries []Entry

func (entries Entries) Scaffold(basePath string) error {
	runner := initEntriesRunner()

	for _, entry := range entries {
		runner.scheduleProcessing(basePath, entry)
	}
	runner.unblockErrorChannelOnCompletion()

	return runner.fetchFirstErrorAndDrainOthers()
}

type entriesRunner struct {
	errChan chan error
	wg      sync.WaitGroup
}

func initEntriesRunner() entriesRunner {
	return entriesRunner{
		errChan: make(chan error),
	}
}

func (runner *entriesRunner) scheduleProcessing(basePath string, entry Entry) {
	runner.wg.Add(1)

	go (func() {
		defer runner.wg.Done()

		if err := entry.Scaffold(basePath); err != nil {
			runner.errChan <- err
		}
	})()
}

func (runner *entriesRunner) unblockErrorChannelOnCompletion() {
	go (func() {
		runner.wg.Wait()

		runner.errChan <- nil
		close(runner.errChan)
	})()
}

func (runner *entriesRunner) fetchFirstErrorAndDrainOthers() error {
	firstErr := <-runner.errChan

	go (func() {
		for range runner.errChan {
		}
	})()
	return firstErr
}
