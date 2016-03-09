# git-grouping
[![Build Status](https://drone.io/github.com/takecy/git-grouping/status.png)](https://drone.io/github.com/takecy/git-grouping/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/takecy/git-grouping)](https://goreportcard.com/report/github.com/takecy/git-grouping)

![](https://img.shields.io/badge/golang-1.6.0-blue.svg?style=flat-square)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/takecy/git-grouping)
![](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)

`git-grouping` is grouping git repositories and execute command to group. 

<br/>
### Motivation 
`microservice` is composed of many of the repository.  
`api` `web` `app` `proxy` `queue` etc...    
I want to manage by grouping the repository to compose the same function.  


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
Print version.
```
$ ggp version
```

Add current repository to group.
```
$ ggp add . apis
```

Remove current repository from group
```
$ ggp del . apis
```

Execute command to gourped repositories.
```
$ ggp exec apis status
$ ggp exec apis fetch --all -p
```

<br/>
### License
[MIT](./LICENSE)
