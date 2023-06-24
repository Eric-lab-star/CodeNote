Plumbing 
	core git

Porcelain
	The prettier user interface on top of core git

# Creating a Git repository

working tree로 사용할 디렉토리만 지정하면 됨

예) 
```
mkdir git-tutorial
cd git-tutorial 
git init
```

위의 코드를 입력하면 .git 폴더가 생서됨

.git 폴더 안에 들어 있는 3가지 주요 파일

HEAD 파일
- ref: refs/heads/main
- 포인터라고 생각하면 됨

objects 폴더
- 내가 만든 프로젝트의 모든 오브젝트를 저장한다
- 오브젝트는 160-bit SHA-1로 식별된다. 이것을 오브젝트 이름이라고 한다.
- 오브젝트에 대한 참조는 40 바이트의 16진법으로 표기된 SHA-1 이름이다.
-  refs의 하위 폴더에는 40바이트 16진법의 이름과 \n 이 포함되어 41바이트의 이름이 저장된다.

refs 폴더
- objects를 참조한다. 
- heads와  tags 라는 폴더를 가지고 있다. 
- heads에는 서로 다른 head가 가리키는 참조를 저장
- tags에는 내가 만든 tag가 저장

# Populating a Git Repository

working tree (working directory)
- hello
- example

위와 같은 파일을 만들었을 때 등록하는 방법
- index 파일을 working tree 상태의 정보로 채우기
- index 파일을 객체로 커밋하기

working tree의 변경 사항을 git에 알려주는 방법
- git update-index
	- 파일의 내용을 깃의 object 형태로 만들어서 저장한다.
	- 업데이트하고 싶은 파일이름을 받는다.
	- --add 플래그를 입력해야 새로운 파일이 등록된다.
	- --remove 플래그를 입력하면 등록된 파일이 사라진다.
	- 명령어를 실행하면 objects 폴더에 새로운 object가 생성된것을 확인 할 수 있다.
	- object 파일의 타입을 읽을려면 `cat-file -t <object name>` 을 통해서 파일 타입을 확인 할 수 있고, 파일의 내용을 읽을려면 `cat-file <type> <object name>` 을 입력하면 된다.
	- object 는 변경할 수 없는 파일이다. 
	- object를 직접적으로 읽을 상황은 거의 없다.
	- update-index 명령어는 index 파일을 생성한다
		- index 파일은 현제 working tree를 설명한다.
	-  현제는 git에 등록을 하지 않았다. 파일에 대하여 설명을 했을 뿐

- git diff-files -p
	- working tree와 index를 비교해서 차이를 보여준다.
	-  git diff 는 단축어


# Committing Git state

index에서 알고 있는 파일을 실제 트리로 커밋을 하기

두 단계로 나누어진다. 

1. tree 객체 만들기
2. tree 객체를 commit 객체로 커밋하기
	1. commit  객체는 tree 정보를 설명한다

- write-tree
	- 현제의 index 상태를 바탕으로 객체를 만든다.
	-  파일이름과 내용을 모아서 깃 디렉토리 객체를 만드는 과정이다 
	- 하지만 이 명령어 혼자 사용되는 경우는 없다. git commit-tree와 같이 사용한다
	- 또한 이 명령어를 혼자서 사용하는 것은 쉽지않다.
- commit-tree
	- 현제 커밋의 부모 커밋을 알아야한다. 
	- 커밋 메시지도 입력을 해야한다.
	- 이 과정을 통해서 .git/refs/heads/main 파일이 생성되고 HEAD파일이 이 위치를 가리키게된다.

위의 모든 과정은 git commit 명령어를 이용해서 편리하게 사용가능하다

# Making a Change

- git diff-index
	- diff-files와 달리 index는 커밋 트리와 index 파일 혹은 working tree 를 비교한다.
	- git diff-index -p HEAD 를 하면 working  tree 와 커밋 트리를 비교한다.
	- 단축어로 diff HEAD 를 하면 된다.
	- index 와 비교를 하고 싶으면  git diff-index --cached HEAD를 하면 된다.
	- 
# Inspecting Change

git diff-tree
- 두 개의 임이의 트리를 받으면 서로 비교를 한다. 
- 한개의 커밋 객체를 받으면 부모와 비교를 한다. 
	- git diff-tree -p HEAD
	- diff-tree --pretty 를 하면 작성자, 날짜, 커밋 메시지가 나온다.

git log -p
- 모든 커밋의 변경 사항을 보여준다.

# Tagging a version

light && annotated tag

light
- git tag  <`tag`> 
	-  .git/refs/tags/ 폴더에 입력한 태그이름의 파일이 생성된다. 
	- 생성된 파일에 HEAD의 내용을 저장한다.

annotated tag
- git tag -s `<tag name>`  
- 실제 git오브젝트로서 태그 이름, 메시지, pgp 사인이 있다. 


# copying repositories

git의 working directory 와 repository는 그냥 폴더이다. 따라서 복사를 해서 이동 시키거나 지울 수 있다. 

 `rm -rf git-tutorial`

git이 있는 폴더 전체를 복사시킬려면 

`cp -a git-tutorial new-git-tutorial`  커맨드를 이용할 수 있다 

복사를 한 이후에는 index 파일을 새로 고침해야한다. 

 `git update-index --refresh`

위의 명령어를 입력하면 index가 최신의 상태로 만들어진다.

# Creating a new branch

깃의 브랜치는 깃  object 데이터베이스를 가리키는 .git/refs이다.

- `git switch -c mybranch`
	- create branch 
	- HEAD 위치를 기준으로 새로운 브랜치를 만든다.
- `git switch -b mybranch earlier-commit`
	-  현제  HEAD가 아닌 과거의 커밋을 기준으로 새로운 브랜치를 만들고 싶으면
 - `git switch main`
	 - 다시 main으로 가고 싶을 때
-  `cat .git/main`
	 - 현제의 브랜치 이름을 확인 하고 싶으면
- git branch
	- 모든 브랜치 이름을 나열할 때
	- `ls .git/refs/heads` 와 동일하다.
-  `git branch <branchname> [startpoint]` 
	- 새로운 브랜치를 만들지만 그 브랜치로 이동하고 싶지 않을 때


# Merging two branches

브랜치를 만드는 이유는 실험을 하기 위함이다. 실험을하고 다시 합치기 위함이다. 

`git commit -m <msg> -i <filename>`
-  -m 플래그는 메시지를 입력하고
- -i 플래그는 index에 추가할 파일이름이다. 

상황 
1. mybranch 에서 파일을 수정함
2. 다른 사람이 main 에서 파일을 수정함

`git merge -m <msg> <branchname>`
merge를 실행한 후에 충돌이 생기면 합병이 안된다. 
합병된 이후에는 gitk 명령어를 이용해서 GUI로 확인 할 수 있다.

- fast-forward merge
	- mybranch를 main으로 합병시킨 다음에 다시 mybranch 로 돌아가서 
	- main을 mybranch에 합병시켜주면 fast-forward merge가 된다.

# Merging external work

외부 저장소에서 파일을 가져오는 명령어는 

`git fetch <remote-repository` 이다.

원격 저장소는 4가지가 있다. 
1. ssh 형식
	- ssh 는 secure shell이다. 
	- 부족한 오브젝트를 발견하고, 최소한의 오브젝트를 교환한다. 
	- 가장 효율적인 방식이다.
	- 이 방식을 사용하기 위해서는 설정하는 과정이 필요하다.
	- [ssh key 생성](https://docs.github.com/ko/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent) 
1. 로컬 패스
	1. ssh 를 이용하는 것은 동일하지만 원격저장소가 아닌 로컬 저장소를 이용한다.
2. git native
3. http

원격 저장소를 가져오면 `merge` 해야한다. 
하지만 fetch 이후에 merge하는 것이 보편적입으로 
`git pull <remote-repository>`  를 이용한다.

 `git config remote.<remotename>.url <url>

 - 위의 명령어는 config 파일에 url을 저장하는 것이다. 
 - config 파일은 .git/config에 존재한다. 
 - 위와 같이 저장을 하면 원격저장소 주소를 반복적으로 사용하는 것 대신에 설정한 이름으로 저장할 수 있다.






