package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
	"github.com/spf13/viper"
)

func AddToSyncQueue(path string) {
	currentUnsynced := viper.GetStringSlice("unsynced")
	currentUnsynced = append(currentUnsynced, path)
	viper.Set("unsynced", currentUnsynced)
	viper.WriteConfig()
}

func GetSyncServerURL() string {
	if viper.GetString("env") == "prod" {
		return "https://ettih.blitzcloud.me/api/admin"
	} else {
		return "http://localhost:3000/api/admin"
	}
}

// func UpdateSyncTimeStamp() error {

// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", GetSyncServerURL()+"/last-sync", nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("admin_token")))
// 	_, err = client.Do(req)
// 	return err
// }

func CloneRepo(path string, URL string) {
	if _, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) {

		err := os.MkdirAll(path, 0766)
		if err != nil {
			log.Fatal(err)
		}
		_, err = git.PlainClone(path, &git.CloneOptions{
			URL:               URL,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

}

func createGitRepoIfNotExists(path string) {
	if _, err := os.Stat(filepath.Join(path, ".git")); os.IsNotExist(err) {
		_, err := git.PlainInit(path, false)
		if err != nil {
			log.Fatalf("failed to initialize git repository: %v", err)
		}
	}

}

func commitNewFilesToGitRepo() {
	repoPath := GetRootDirectory() + "/local"
	createGitRepoIfNotExists(repoPath)
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("failed to open repository: %v", err)
	}

	// Get the working tree
	worktree, err := repo.Worktree()
	if err != nil {
		log.Fatalf("failed to get worktree: %v", err)
	}
	_, err = worktree.Add(".")
	if err != nil {
		log.Fatalf("failed to add all files: %v", err)
	}
	// fmt.Println("All changes added to staging area.")
	// fmt.Println("Performing 'git commit -m \"Added new files\"'...")
	_, err = worktree.Commit("Added new lab", &git.CommitOptions{
		Author: &object.Signature{
			Name:  viper.GetString("git_name"),
			Email: viper.GetString("git_email"),
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatalf("failed to commit changes: %v", err)
	}

	// fmt.Printf("Successfully committed with hash: %s\n", commitHash)
}

func addPathToRootGitIgnore(path string) {}

func UpdateAndPushGitRepo(repoPath string) error {

	auth := &http.BasicAuth{
		Username: viper.GetString("git_name"),
		Password: viper.GetString("git_token"),
	}
	// Open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return fmt.Errorf("failed to open repository at %s: %w", repoPath, err)
	}

	// Get the worktree to interact with the repository's files
	worktree, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	_, err = worktree.Add(".")
	if err != nil {
		log.Fatalf("failed to add all files: %v", err)
	}
	// 1. Pull the latest changes from the remote
	fmt.Println("Pulling latest changes...")
	err = worktree.Pull(&git.PullOptions{
		RemoteName: "origin", // Assuming "origin" is your remote name
		Auth:       auth,
	})
	if err != nil && err != git.NoErrAlreadyUpToDate {
		return fmt.Errorf("failed to pull latest changes: %w", err)
	}
	if err == git.NoErrAlreadyUpToDate {
		fmt.Println("Repository is already up to date.")
	} else {
		fmt.Println("Pull complete.")
	}

	// Get the status of the repository to find added files
	status, err := worktree.Status()
	if err != nil {
		return fmt.Errorf("failed to get worktree status: %w", err)
	}

	hasChangesToCommit := false
	for _, s := range status {
		if s.Staging == git.Added || s.Staging == git.Modified || s.Staging == git.Deleted {
			hasChangesToCommit = true
			break
		}
	}

	if !hasChangesToCommit {
		fmt.Println("No new files or changes to commit.")
	} else {
		// Add all changes to the staging area
		fmt.Println("Adding all changes to staging...")
		err = worktree.AddWithOptions(&git.AddOptions{
			All: true, // Add all changed files
		})
		if err != nil {
			return fmt.Errorf("failed to add all changes: %w", err)
		}

		// 2. Commit the added files with the specified message
		fmt.Println("Committing changes...")
		commitMessage := "Update: " + GetPrettyDate(time.Now()) + " Added new lab"
		_, err = worktree.Commit(commitMessage, &git.CommitOptions{
			Author: &object.Signature{
				Name:  viper.GetString("git_name"),
				Email: viper.GetString("git_email"),
				When:  time.Now(),
			},
		})
		if err != nil {
			return fmt.Errorf("failed to commit changes: %w", err)
		}
		fmt.Printf("Commit '%s' created successfully.\n", commitMessage)

		// 3. Push the changes
		fmt.Println("Pushing changes...")
	}
	err = repo.Push(&git.PushOptions{
		Auth: auth,
	})
	if err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}
	fmt.Println("Push complete.")

	return nil
}
