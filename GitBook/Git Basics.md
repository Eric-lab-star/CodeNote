# Git Basics

## Getting a Git Repository

- Initializing a Repository in an Existing Directory
  1. `mkdir <dirname>`
  2. `git init`

## Recording Changes to the Repository

### Checking the Status of Your Files

`git status`

- difference between index file and current HEAD commit --> 커밋을 하지 않음
- working tree and index file --> 파일이 수정되었음
- working tree note tracked by git --> 새로운 파일이 생성됨

### `git add` command

1. tracking new files
2. staging modified files
3. marking merge-confict files as solved

`git add` is multipurpose command

If you modified file after `git add`, you need to run `git add` again. Git consider two files as different version

## Short Status

`git status -s`

Provides more simplified version of `git status`.
New files that have been added to the staging area have A.
Modified files have M.

## Ignoring Files

The rules for the patterns

- Blank lines or lines starting with `#` are ignored
- matches zero or more characters
- `[abc]` matches any character inside the brackets
- `?` matches a single character
- `[0-9]` matches any characters between them
- `a/**/b` matches nested directories

```zsh
# ignore all .a files
*.a

# don't ignore lib.a
!lib.a

# only ignore TODO file in the current directory, not subdir /TODO
/TODO

# ignore doc/notes.txt, but not doc/server/arc.txt
doc/*.txt

#ignore all files in any directory named build
build/

# ignore all .pdf files ind the doc/ directory and any of its subdirectories
doc/**/*.pdf

```

## Viewing Your Staged and Unstaged Changes

if `git status` is too vague, use `git diff`.
`git diff` is used to provide answers to following questions.

1. What have you changed but not yet staged?
2. What have you staged that are about to commit?

`git diff`

- compares what is in your working directory with what is in your staging area
- staging area<---> working tree
- unstaged <--> staged

`git diff --cached`

- index <---> latest commit
- waiting to be committed <---> latest commit

git status 는 파일이름을 보여주지만 git diff는 파일의 줄 까지 보여준다.

## Committing Your Changes

`git commit -v`

- To check diff of your change add -v option
- diff will be striped off

## Skipping the Staging Area

`git commit -a `

- git automatically stage every file that is already trakced before.

`git commit -i <파일 이름>`

- 추가할 파일 이름을 선택할 수 있다.

## Removing Files

`git rm -f <filename>`

- index 파일에서 파일을 지우고, working directory 에서도 지운다.

`git rm --cached <filename>`

- index 파일에서만 파일을 지운다.
- this is usefull when you forgot to add file to .gitignore.

## Moving Files

`git mv file_from file_to`

- git does not explictly track file movement. Instead it implicitly track file movement.

## Viewing the commit History

`git log -p -1`
show diff

`git log --stat`
abbreviates diffs

`git log --pretty=oneline`
list commit name and commit message of all commits

`git log --graph`
show little nice graph showing your branch and merge history

## Undoing Things

`git commit --amend`

- 이미 작성한 커밋을 수정하고 싶을 때 사용한다.
- 커밋 메시지를 수정할 수 있고 파일을 추가할 수 있다.

## Unstaging a Staged File

만약 어떻게 햐야 할 지 방법을 모를때는 `git status` 을 잘 읽어 보면 도움이 된다.
`git reset HEAD <filename>`

- reset을 이용해서 git add 를 취소한다.

## Unmodifying a Modified File

파일을 수정을 했으나 수정하기 전으로 만들고 싶으면
`git checkout -- <filename>`

주의 사항

- anything that is commited can be recovered, but uncommitted file cannot be recoverd
- `git checkout -- <file>`은 위험하다. 이 실행은 되돌릴수 없기 때문이다.
- 확실한 경우가 아니라면 사용하지 않은 편이 좋다.

최신 깃에서는 reset 대신 restore 이라는 이름을 사용한다.

undo modifired working directory
`git restore <filename>`

`git add` 취소
`git restore --staged <filename>`

## Working with Remotes

## Showing Your Remotes

`git remote`

- list the the shortnames of each remote
- `origin` is default remote branchname created by `git clone` command

`git clone`

- git clone sets remote branch

## Adding Remote Repositories

`git clone`

- implicitly add `origin` remote

`git remote add <shortname> <url>`

- add remote repository explicitly

`git fetch <shortname>`

- `shortname` 에 있는 데이터를 가져온다.
- shortname에 있는 데이터는 remote 브렌치에 있다. local 브렌치로 가져오기 위해서는 merge 해야한다.
  - `git branch -r` --> 브랜치 이름 확인
  - `git merge <remote branch>`

If remote branch and local branch does not share same commit history, git refuses to merge two branches

## Fetching and Pulling from your remotes

To get data from remote proj.
`git fetch <remote>`

- `git fetch` download data to your local repository
- it doesn't automatically merge it with any of your work
- you have to merge it manually

## Pushing to Your Remotes

`git push origin master`

- if you and someone else clone at the same time and they push upstream and then you stream, your push will reightly be rejected you will have to fetch work first.

## Inspecting a Remote

`git remote show origin`

## Renaming and Removing Remotes

`git remote rename <old> <new>`

`git remote remove <remote>`

- remve remove branch

## Tagging

Git has the ability to tag specific points in a repository's
history as being important. Typically, people use this
functionaliy to mark releas points.

## List Your Tags

`git tag`

- match pattern
  - `git tag -l "v1.8.5"`

## Creating Tags

git supports two types of tags: _lightweight_ and **annotated**

lightweight

- branch that doesn't change

annotated

- stored as full objects in the git database
- contains
  - tagger name
  - email
  - date
  - tagging message

## Annotated Tags

`git tag -a v1.4 -m "my version 1.4"`
`git show v1.4`

## Lightweight Tags

commit checksum stored in file - no other information is kept.
to provide lightweight tag just provide a tag name.

## Tagging Later

If you want to make a tag from previously commited file, append commit name.

`git tag -a v1.2 9fceb02`
나중에 git log를 통해서 commit 이름을 파악한 다음 태그를 지정할 수 있다.

## Sharing Tags

`git push origin <tagname>`

if you have a lot of tags that you want to push at once, you can also use the --tags options to the git push command.
`git push origin --tags`

## Deleting Tags

`git tag -d v1.4`

deleting tag from a remote server
`git push <remote> :refs/tags/<tagname>`
`git push origin --delete <tagname>`
둘 중에 하나만 사용하면 된다.

## Checking out tags

`git checkout <tagname>`

> [!note]- Note
> git checkout `<tagname>` puts your repository in "detached head" state. In "detached HEAD" state, if you make changes and then created a commit, the tag will stay the same, but your new commit wont belong to any branch and will be unreachable, except by the exact commit hash. **Thus, if you need to make changes - you will generally want to create a branch.**

`git checkout -b <branchname> <tagname>`

## Git Aliases

How to set up git Aliases

`git config --global alias.co checkout`
