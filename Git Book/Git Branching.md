# Branches in a Nutshell
resoure: [git branching](https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell)
## Branches in a Nutshell

When you make a commit, Git stores a commit object that contains a pointer to the snapshot of the content you staged.


Staging the files computes a checksum for each one, stores that version of the file in the Git repository, and add that checksum to the staging area

![[commit-and-tree.png]]

| Term   | Definition                                                                                  |
| ------ |:------------------------------------------------------------------------------------------- |
| blob   | represents version of files                                                                 |
| tree   | A snapshotlists the content of the directory and specifies which file names are stored as which blobs |
| commit | pointing to root tree                                                                       |

![[commits-and-parents.png]]


if you make some change and commit again, next commit stores a pointer to parent commit that came immeditely before it.

A branch in Git is simply a light weight movable pointer to one of these commits. As you start making commits, you're given a master branch that points to the last commit you made. Everytime you commit the mater branch moves forward automatically.

![[branch-and-history.png]]


## Creating a New Branch

When you run command `git branch testing`, this creates new branch testing pointing to the same commit.

