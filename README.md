# git-wta

Add a worktree with opinionated naming convention

It will create new worktree from given commit-ish. The worktree path will be:
    <the main worktree path>+<given commit-ish>

If the main worktree is a bare repository, .git suffix in the path will be removed.

## Usage

```
$ git wta <commit-ish>
```

## Installation

```
$ go install github.com/yoichi/git-wta@latest
```

## License

MIT

## Author

Yoichi NAKAYAMA