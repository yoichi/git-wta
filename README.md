# git-wta

Add a worktree with opinionated naming convention

It will create new worktree from given commit-ish. The worktree path will be:
```
<the main worktree path>+<given commit-ish>
```

If the main worktree is a bare repository, .git suffix in the path will be removed.

## Usage

```
$ git wta <commit-ish>
```

## Installation

```
$ go install github.com/yoichi/git-wta@latest
```

## Shell completion

To enable bash/zsh completion, add the following to your shell settings:

```
_git_wta ()
{
        local dwim_opt="$(__git_checkout_default_dwim_mode)"
        __git_complete_refs $dwim_opt --mode="refs"
}
```

## License

MIT

## Author

Yoichi NAKAYAMA