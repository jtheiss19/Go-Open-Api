# Go-Open-Api
<div align="center">
  
![GitHub contributors](https://img.shields.io/github/contributors/jtheiss19/MARI-Engine)
![GitHub forks](https://img.shields.io/github/forks/jtheiss19/MARI-Engine?label=Forks)
![GitHub stars](https://img.shields.io/github/stars/jtheiss19/MARI-Engine?style=Stars)
![GitHub issues](https://img.shields.io/github/issues-raw/jtheiss19/MARI-Engine)
[![Go Report Card](https://goreportcard.com/badge/github.com/jtheiss19/MARI-Engine)](https://goreportcard.com/report/github.com/jtheiss19/MARI-Engine)

</div>

<a href="https://github.com/jtheiss19/Go-Open-API"><img src="./sample.png" alt="Demo Picture"></a>


## Table of Contents

- [Introduction](#Introduction)
- [Installation](#installation)
- [Features](#features)
- [License](#license)



# Introduction

I work with gateways often in my industry. Time and time again I found myself frustrated with how modern gateways work and their instalation process. Modern gateways also typically lack support for new tools such as kubernetes and docker. 

Go-Open-API is my solution to these problems. I wanted an API gateway that is easily configurable and easy to work with. To this end I plan on adding a docker image, easy steps for building your own image, support for working withing a kube cluster, cluster of its own for easy settings replication.

Some of this has already happened. For example you can currently (with a modern terminal) ssh into a machine and access the terminal based GUI and run commands and check logs without third party support or additional overhead.

# Installation

> go get github.com/jtheiss19/Go-Open-Api


## Setup

> Navigate to the ui folder and run the main package

```shell
$ cd $GOPATH/src/github.com/jtheiss19/Go-Open-Api
$ go run .
```

# Features
* Mouse Support [DONE]
* Keyboard Support
* Custom HTTP Sorting
* TLS Support

# License

**[MIT license](http://opensource.org/licenses/mit-license.php)**
