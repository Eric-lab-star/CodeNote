---
date created: 2023-06-01 14:44
date updated: 2023-06-01 14:46
title: Getting Stated
tag: git
---

# Getting Started

## About Version Control

Version control is a system that records changes to a file or set of files over time so that you can recall specific versions later.

Version control gives an opportunity to revert file back to previous state.

## What is Git?

### Snapshots, Not Differences

Git thinks of its data more like a series of snapshots of a miniature filesystem. Git takes a picture of what all your files look like at that moment and stores a reference to that snapshot.

To be efficient, if files have not changed, Git doesn't store the file again, just a link to the previous identical file it has already stored.

### Git Has Integrity

Everything in Git is checksummed before it is stored and is then referred to by that checksum. You can't lose information in transit or get file corruption without Git being able to detect it.

The mechanism that Git uses for this checksumming is called a SHA-1 has. This is a 40-character string composed of hexadecimal characters and calculated based on the content of a file or directory structure in Git.

Git stores everything in its database not by file name but by the hash values of its contents.

### Git Generally Only Adds Data

When you do actions in Git, nearly all of them only add data to the Git database. It is hard to get the system to do anything that is not undoabnle or to make it erase data in any way.

### The Three States

Files inside Git has three main states.

1. modified
2. staged
3. committed

- modified means that you have changed file but have not committed it to your database yet

- staged meas that you have marked a modified file in its current version to go into your next commit

- committed means that the data is safely stored in you local database

Git has three sections

1. Working directory
2. Staging Area
3. Git directory

Working tree is a single checkout of one version of the project.

Staging area is a file, contained in your Git directory, that stores information about what will go into your next commit. It is also called as "index"

Git directory is where Git stores the metadata and object database for your project. It is what is copied when you clone a repository from another computer.

## Setup

- config

  - git confit --list --show-origin

- Your Identity

  - Every Git commit uses this information, ad it's immutably baked into the commits you started creating
  - `git config --global user.name "<name>"`
  - `git config --global user.email <email>`

- Your Editor

  - `git config --global core.editor <editor>`

- Default Branch name

  - `git config --global init.defaultBranch <name>`

- Checking Setting
  - `git config --list`
