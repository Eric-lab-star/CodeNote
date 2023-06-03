# Git Branching

## Branches in a Nutshell

A branch in Git is simply movable pointer to one of these commits.

## Creating a New Branches

`git branch testing`

HEAD is special pointer pointing current location in local branch

## Switching Branches

`git checkout <branch name>`

you can use `git switch`
This moves HEAD to point to the new branch given in the command

`git log --oneline --graph --all`

This prompt shows branched in simple graph

## Basic Branching

`git checkout -b iss53`

this is shorthand for:

`git branch iss53`
`git checkout iss53`

### fast-forward

when you try to merge one commit with a commit that can be reached by following the first commit's history, Git simplifies things by moving the pointer forward because there is no divergent work to merge together.

### delte branch

`git branch -d hotfix`

after merge branch can be deleted

## Basic Merge Conflicts

If you changed the same part of the same file differently in the two branches you're merging Git won't be able to merge them cleanly.

## Branch mamangement

`git branch`

show lits of branches

`git branch -v`

show list of branches and last commit from each branch

`git branch --merged`

list it this branchs are merged branch. branch without \* is ok to delete

`git branch --no-merged`

lisf of branched that are not merged. trying to delete any branch on this lits will be rejected. if you really want to delete, use `git branch -D <branch name>`

## Changing a branch name

`git branch --move <bad branch name> <corrected branch name>`

## rename branch locally

`git push --set-upstream origin <corrected branch name>`

to let others see the corrected branch name, but this command does not delete <bad branch name>

## delete branch on remote server

`git push origin --delete <bad-branch-name>`

delete bad branch name in remote server

## Long Running Branches

most developers use two branches. One branch to keep track of stable code, an another branch to do test and development before release.
Having long-running branched isn't necessary but it is often helpful, especailly when you are dealing with very large or complex projects.

## Topic Branches

A topic branch is a short-lived branch that you created and use for a single particular feature or related work.

## Remote Branches

### Remote-tracking branches

remote tracking branch names take the form <remote>/<branch>.

### syschronize with remote branch

`git fetch <remote>`

this command moves origin/main branch, but does not merge

## Pushing

`git push <remote> <branch>`

to share branch with others

`git push <remote> <branch>:<newname>`

if you want to use other name on the remote server

## Tracking Branches

Tracking branch are local branches that have a direct relationship to a remote branch. **If you are on a tracking branch and type 'git pull', Git automatically knows which server to fetch from and which branch to merge in**.

When you prompt `git clone`, main branch is setup as tracking branch that tracks origin/main.
origin/main branch becomes upstream branch.

`git checkout --track <remote>/<branch>`
`git checkout -b <branch> <remote>/<branch>`

This command set another tracking branch. It can be ones that track branches on other remotes, or don't track master branch.
To use this command, you should have <remote>/<branch> previously and you should not have same branch name.

If you want to use local branch to track remote branch, use `git branch -u <origin>/<branch>` or `git branch --set-upstream-to=origin/main main`

`git branch -u origin/main`
if you want to set upstream branch.

`git branch -vv`
to list tracking branches that have been set up.

## Pulling

If you have tracking branch `git pull` will automatically fetch and pull data from upstream branch

`git pull`

## Deleting Remote Branches

`git push origin --delete <branch>`

to delete branch from server

## The basic Rebase

with rebase, you can take all the changes that were commitied on one branch and replay them on a differenct branch.

if you are on dev branch
`git rebase main` moves current commit set main branch as base of branch

after rebase dev branch, you can checkout main branch to fast-forward to dev branch.

The out of history using rebase gives more cleaner output than using merge.
rebasing reduces work for project maintainer

## More Interesting Rebase

`git rebase --onto master server client`

## The perils of Rebasing

Do not rebase commits that exist ouside your repository and that people may have based work on.
