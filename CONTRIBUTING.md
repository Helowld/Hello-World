Contributing to Hello-World
=======================

Welcome
-------

We invite you to join our team! Everyone is welcome to contribute code
via pull requests, to file issues on GitHub.

## Guidelines for Merging

- while giving the Pull requests please follow the template [PULL_REQUEST_TEMPLATE.md](./PULL_REQUEST_TEMPLATE.md)
- fileName should be hello_world.extension like hello_world.c(extension should be in lowercase)
- you should consider the reader as newbies, so explain the code according to that
- the program has to contain the comments which has to explain the program precisely, for example check this out [hello_world.c](https://github.com/rustiever/Hello-World/blob/main/hello_world.c).
- add respective language's official website link and resource material to get started.(For some language official websites or resource materials won't be availabale, in that case add a popular website which has resource material to respective language)
- PR which will be having the specific file name as given in the issue, the one whose commit descriptions are relevant, and the code is readable, would be merged. 

For docs, make sure that you use markdown syntax. Checkout these links: [Markdown Guide](https://www.markdownguide.org/), [Markdown Cheat-Sheet](https://www.markdownguide.org/cheat-sheet/).

## How to Collaborate:

1. Fork the repository to your own GitHub account.

2. Clone the repository to your local machine
```
$ git clone "https://www.github.com/{Username}/Hello-World.git"
```
where username is your GitHub account username.

3. Create a branch where you can do your local work.
Never work on **master** branch as we do not allow master commits except by admins.
```
$ git branch {branchname}
$ git checkout branchname
```

4. Do your work and stage your changes.
```
$ git add <filename>
```

5. Commit you changes with a commit message containing your name, file(s) worked upon, changes added.

6. Push changes to your forked repository
```
$ git push -u origin branchname
```
7. Create a pull request to the upstream repository.

## Synchronize forked repository with Upstream repository

1. Create upstream as our repository
```
$ git remote add upstream "https://www.github.com/rustiever/Hello-World"
```

2. Fetch upstream changes in local machine
```
$ git fetch upstream
```

3. Switch to master branch
```
$ git checkout master
```

4. Merge changes in local machine
```
$ git merge upstream/master
```

5. Push changes to your forked GitHub repository
```
$ git push -f origin master
```
