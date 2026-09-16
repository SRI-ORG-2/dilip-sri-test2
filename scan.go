package main

import "fmt"

func triggerScan(aq AquaClient, repo Repository, aquaTriggerScanBranch string) {
	result, err := aq.TriggerRepository5canV2Detailed(repo.AquasecRepoID, aquaTriggerScanBranch)

	if err != nil {
		if result != nil {
			fmt.Printf("Failed to trigger scan for %s (%s): %v [status=%d queued=%t body=%q]\n", repo.FullName, repo.AquasecRepoID, err, result.StatusCode, result.Queued, result.ResponseBody)
			return
		}

		fmt.Printf("Failed to trigger scan for %s (%s): %v\n", repo.FullName, repo.AquasecRepoID, err)
		return
	}

	fmt.Printf("Triggered scan for %s (%s) [status=%d queued=%t body=%q]\n", repo.FullName, repo.AquasecRepoID, result.StatusCode, result.Queued, result.ResponseBody)
}
