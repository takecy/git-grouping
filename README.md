# ggp

`ggp` is **g**it **g**rou**p**ing.  
To group multiple Git repositories.  
For the group, You can Git command is run.  


[![Build Status](https://drone.io/github.com/takecy/git-grouping/status.png)](https://drone.io/github.com/takecy/git-grouping/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/takecy/git-grouping)](https://goreportcard.com/report/github.com/takecy/git-grouping)
![](https://img.shields.io/badge/golang-1.6.0-blue.svg?style=flat-square)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/takecy/git-grouping)
![](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)

<br/>
### Motivation
The microservice is composed by many of the repository.  
`api` `web` `sp_app` `proxy` `queue` etc...    
also `api` component is composed by many of repository.   
`core` `front` `batch` `cache` etc...  
also `core` component is ///

I want to manage by grouping the repository to compose the same function or component.  


<br/>
### Install

#### via Go
```
$ go get github.com/takecy/git-grouping/ggp
```

#### via Binary
[Download](https://github.com/takecy/git-grouping/releases) and Copy to your `$PATH`.


### UnInstall
Delete your `ggp` binary.  
And delete `$HOME/.ggp` directory.  


<br/>
### Usage
All usages
```
$ ggp
```
```
add              Add repository to group
git              Execute git commands to group of repository
info             Show configuration.
ls               List all group.
rm               Remove repository form group
rmg              Remove group
version          Print version
```
Show detail help.
```
$ ggp help add
add <group name> <dir path>
```

<br/>
#### Example
Show all groups.
```
$ ggp ls
```
Show group info.
```
$ ggp info apis
```

Add current repository to group.
```
$ ggp add apis .
```
Add specific repository to group.
```
$ ggp add /git/your/cool-tools
```

Remove current repository from group
```
$ ggp rm apis .
```

Remove group.
```
$ ggp rmg apis
```

Execute command to gourped repositories.
```
$ ggp git --gg apis status
$ ggp git --gg apis fetch --all -p
```

<br/>
### Development
Vendoring by [godep](https://github.com/tools/godep).
```
$ git clone git@github.com:takecy/git-grouping.git
$ cd git-grouping
$ make install_tools
$ godep restore
$ DEBUG=* go run ggp/main.go
```


<br/>
### License
[MIT](./LICENSE)
