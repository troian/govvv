package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const versionFile = "VERSION"

// GetFlags collects data to be passed as ldflags.
func GetFlags(dir string, args []string) (map[string]string, error) {
	repo := git{dir}
	gitBranch := repo.Branch()
	gitCommit, err := repo.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to get commit: %v", err)
	}
	gitCommitFull, err := repo.CommitFull()
	if err != nil {
		return nil, fmt.Errorf("failed to get full commit: %v", err)
	}
	gitCommitMsg, err := repo.CommitMsg()
	if err != nil {
		return nil, fmt.Errorf("failed to get commit message: %v", err)
	}
	gitState, err := repo.State()
	if err != nil {
		return nil, fmt.Errorf("failed to get repository state: %v", err)
	}
	gitSummary, err := repo.Summary()
	if err != nil {
		return nil, fmt.Errorf("failed to get repository summary: %v", err)
	}

	// prefix keys with package to be used by ldflags -X
	pkg := defaultPackage
	if value, ok := collectGovvvDirective(args, flPackage); ok {
		pkg = value
	}

	gitCommitMsg = strings.Replace(gitCommitMsg, "'", "\\'", -1)
	gitCommitMsg = strings.Replace(gitCommitMsg, "-", "\\-", -1)

	v := map[string]string{
		pkg + ".BuildDate":     date(),
		pkg + ".GitCommit":     gitCommit,
		pkg + ".GitCommitFull": gitCommitFull,
		pkg + ".GitCommitMsg":  gitCommitMsg,
		pkg + ".GitBranch":     gitBranch,
		pkg + ".GitState":      gitState,
		pkg + ".GitSummary":    gitSummary,
	}

	// calculate the version
	if value, ok := collectGovvvDirective(args, flVersion); ok {
		v[pkg+".Version"] = value
	} else if value, err = versionFromFile(dir); err == nil {
		v[pkg+".Version"] = value
	} else if value, err = versionFromGit(repo); err == nil {
		v[pkg+".Version"] = value
	} else {
		return nil, fmt.Errorf("failed to get repository version: %v", err)
	}

	return v, nil
}

// date returns the UTC date formatted in RFC 3339 layout.
func date() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// versionFromFile looks for a file named VERSION in dir if it exists and
// returns its contents by trimming the whitespace around it. If the file
// does not exist, it does not return any errors
func versionFromFile(dir string) (string, error) {
	fp := filepath.Join(dir, versionFile)
	b, err := ioutil.ReadFile(fp)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("failed to find version file %s: %v", fp, err)
	}

	if err != nil {
		return "", fmt.Errorf("failed to read version file %s: %v", fp, err)
	}
	return string(bytes.TrimSpace(b)), nil
}

func versionFromGit(repo git) (string, error) {
	tag, err := repo.Tag()
	if err != nil {
		return "", err
	}

	var cmCount string
	cmCount, err = repo.CommitsCount(tag)
	if err != nil {
		return "", err
	}

	var cmInt int
	if cmInt, err = strconv.Atoi(cmCount); err != nil {
		return "", err
	}

	if cmInt > 0 {
		tag += "-rc." + cmCount
	}

	return tag, nil
}
