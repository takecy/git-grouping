# git-grouping
[![Build Status](https://drone.io/github.com/takecy/git-grouping/status.png)](https://drone.io/github.com/takecy/git-grouping/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/takecy/git-grouping)](https://goreportcard.com/report/github.com/takecy/git-grouping)

![](https://img.shields.io/badge/golang-1.6.0-blue.svg?style=flat-square)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/takecy/git-grouping)
![](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)

`git-grouping` is grouping git repositories and execute command to group. 

<br/>
### Motivation 
The microservice is composed by many of the repository.  
`api` `web` `sp_app` `proxy` `queue` etc...    
also `api` component is composed by many of repository.   
`core` `fornt` `batch` `cache` etc...  
also `core` component is ///

I want to manage by grouping the repository to compose the same function or component.  


<br/>
### Install

#### via Go
```
$ go get github.com/takecy/git-grouping/...
```

#### via Binary
[Download](https://github.com/takecy/git-grouping/releases) and Copy to your `$PATH`.


<br/>
### Usage
All usages 
```
$ ggp
```

Print version. [latest is here.](https://github.com/takecy/git-grouping/releases)
```
$ ggp version
```

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
Add specifi repository to group.
```
$ ggp add /git/your/cool-tools
```

Remove current repository from group
```
$ ggp rm apis .
```

Execute command to gourped repositories.
```
$ ggp exec apis status
$ ggp exec apis fetch --all -p
```

<br/>
### Development
Vendoring by [godep](https://github.com/tools/godep).
```
$ git clone git@github.com:takecy/git-grouping.git
$ cd git-grouping
$ make install_tools
$ godep restore
$ go run ggp/main.go
```


<br/>
### License
[MIT](./LICENSE)
