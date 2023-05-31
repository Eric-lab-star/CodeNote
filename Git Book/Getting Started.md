
## About Version Control

버전 관리를 통해서 실수를 되돌릴 수 있게 할 수 있다.

# Setup

- config
	- git confit --list --show-origin
		- view all of your settings and where they are coming from
	- .git/config
		- config file in the git repo of whatever you're using. Specific to single repository
	- ~/.gitconfig || ~/.config/git/config
		-  Value specific personally to you, the user.
		- git conifg --global
			- read and write to this file

- Your Identity
	- Every Git commit uses this information, ad it's immutably baked into the commits you started createing
	- `git config --global user.name "<name>"`
	- `git config --global user.email <email>`
 
- Your Editor
	- `git config --global core.editor <editor>`
 
- Default Branch name
	- `git config --global init.defaultBranch <name>`

- Checking Setting
	- `git config --list`

