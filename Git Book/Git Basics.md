# Getting a Git Repository

- Initializing a Repository in an Existing Directory
	1. `mkdir <dirname>` 
	2. `git init`


# Recording Changes to the Repository

## Checking the Status of Your Files
` git status`

- difference between index file and current HEAD commit --> 커밋을 하지 않음
- working tree and index file --> 파일이 수정되었음
- working tree note tracked by git --> 새로운 파일이 생성됨
- 

## Ignoring Files

The rules for the patterns

- Black lines or lines starting with `#` are ignored
- * matches zero or more characters
- `[abc]` matches any character inside the brackets
- `?` matches a single character
- `[0-9]` 0에서 9 범위
```
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

`git diff`
- index  <---> working tree

`git diff --cached`
- index <---> latest commit

git status 는 파일이름을 보여주지만 git diff는 파일의 줄 까지 보여준다.

## Committing Your Changes
`git commit -a ` 
- index 에 저장된 모든 파일을 업데이트하고 커밋

`git commit -i <파일 이름>`
- 추가할 파일 이름을 선택할 수 있다. 

## Removing Files
 `git rm <filename>` 
 - index 파일에서 파일을 지우고, working directory 에서도 지운다. 

 `git rm --cached <filename>`
 - index 파일에서만 파일을 지운다.

## Moving Files

`git mv file_from file_to`


# Viewing the commit History

`git log -p -1`
커밋 기록을 가장 최신 커밋부터 보여준다. -p 는  상세 정보를 보여주고 -1은 보여줄 기록의 갯수이다. 

`git log --stat`
커밋 기록을 간략하게 보여준다. 

## Undoing Things

`git commit --amend`
- 이미 작성한 커밋을 수정하고 싶을 때 사용한다. 
- 커밋 메시지를 수정할 수 있고 파일을 추가할 수 있다.

## Unstaging a Staged File

만약 어떻게 햐야 할 지 방법을 모를때는  `git status` 을 잘 읽어 보면 도움이 된다.

`git reset HEAD <filename>
- reset을 이용해서 git add 를 취소한다. 

## Unmodifying a Modified File

파일을 수정을 했으나 수정하기 전으로 만들고 싶으면 
`git checkout -- <filename>`

주의 사항 
- `git checkout -- <file>`은 위험하다. 이 실행은 되돌릴수 없기 때문이다. 
- 확실한 경우가 아니라면 사용하지 않은 편이 좋다. 

커밋을 했다면 `--amend` 를 이용해서 수정을 할 수 있지만  `checkout `을 이용한 수정을 되돌 릴 수 없다.

최신 깃에서는 reset 대신 restore 이라는 이름을 사용한다. 

파일 수정 취소
 `git restore <filename>`

`git add ` 취소 
 `git restore --staged <filename>`


# Working with Remotes


## Showing Your Remotes
`git remote`
- list the the shortnames of each remote
-  `origin` is default name git give

`git clone`
-  git clone sets remote branch

## Adding Remote Repositories

`git clone`
- implicitly add `origin` remote

`git remote add <shortname> <url>`
- add remote repository

`git fetch <shortname>`
-  `shortname` 에 있는 데이터를 가져온다.
- shortname에 있는 데이터는 remote 브렌치에 있다. local 브렌치로 가져오기 위해서는 merge 해야한다. 
	- `git branch -r` --> 브랜치 이름 확인
	-  `git merge <remote branch>` 


## Fetching and Pulling from your remotes

To get data from remote proj.
`git fetch <remote>`
- `git fetch` download data to your local repository
- it doesn't automatically merge it with any of your work
- you have to merge it manually

## Pushing to Your Remotes

`git push origin master`
- if you and someone else clone at the same time and they push upstream and then you stream, your push will reightly be rejected you will have to fetch work first.

## Renaming and Removing Remotes
`git remote rename <old> <new>`

# Tagging

Git has the ability to tag specific points in a repository's 
history as being important. Typically, people use this 
functionaliy to mark releas points.

## List Your Tags
`git tag`
- match pattern
	- `git tag -l "v1.8.5"`

## Creating Tags

git supports two types of tags: *lightweight* and **annotated**

lightweight
-  branch that doesn't change

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
> git checkout `<tagname>` puts your repository in "detached head" state. In "detached HEAD" state, if you make changes and then created a commit, the tag will stay the same, but your new commit wont belong to any branch and will be unreachable, except by the exact commit hash. Thus, if you need to make changes - you will generally want to create a branch. 

`git checkout -b <branchname> <tagname>`





