package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type WorkingTree struct {
	fullPath string
	isBare   bool
}

func getMainWorkingTree() (WorkingTree, error) {
	ls, err := exec.Command("git", "worktree", "list", "--porcelain", "-z").Output()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			fmt.Printf("%s", exitErr.Stderr)
		}
		fmt.Println(err.Error())
		return WorkingTree{"", false}, err
	}
	delimiter := []byte{0}
	parts := bytes.Split(ls, delimiter)
	worktree := string(parts[0])
	prefix := "worktree "
	if !strings.HasPrefix(worktree, prefix) {
		err = fmt.Errorf("invalid output from git worktree: '%s'", worktree)
		fmt.Println(err.Error())
		return WorkingTree{"", false}, err
	}
	return WorkingTree{worktree[len(prefix):], string(parts[1]) == "bare"}, nil
}

func createNewWorkingTree(mainWorkingTree WorkingTree, ref string) error {
	base := mainWorkingTree.fullPath
	gitSuffix := ".git"
	if mainWorkingTree.isBare && strings.HasSuffix(base, gitSuffix) {
		base = base[:len(base)-len(gitSuffix)]
	}
	new := base + "+" + ref
	add, err := exec.Command("git", "worktree", "add", "--checkout", new, ref).Output()
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok {
			fmt.Printf("%s", exitErr.Stderr)
		}
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("Added worktree %s (checking out '%s')\n", new, ref)
	fmt.Print(string(add))
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: git-wta <commit-ish>")
		os.Exit(1)
	}
	ref := os.Args[1]
	mainWorkingTree, err := getMainWorkingTree()
	if err != nil {
		os.Exit(1)
	}
	err = createNewWorkingTree(mainWorkingTree, ref)
	if err != nil {
		os.Exit(1)
	}
}
