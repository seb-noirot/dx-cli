## DX CLI

After watching the inspiring [talk at NDC Porto](https://ndcporto.com/agenda/a-love-letter-to-long-lines-and-other-dx-stories-0lpb/0cvazslx2b4) of [Jo Franchetti](https://github.com/thisisjofrank), I wanted to improve our onboarding experience.
She mentioned that README explaining how to setup projects are a great plus!

But in our case, the biggest trouble one can get ( and also if you spoil coffee on your laptop and need to start anew), is to setup your environment with gitlab, with docker, with our kubernetes and so on.

As no team really own that part of the process, we normally send people to the official documentation.

From this situation, I first though to create just a documentation but then, I though, why not doing a CLI? Why not learning Go in the process??

And here I am, creating a go cli to setup our environment!

Hope you will enjoy it!

### Setup

First you need to install go (and I hope it is the last thing you need to install manually)
Look at the process on their [page](https://go.dev/doc/install) 

```go
go run main.go
```
And it will list you the command that the cli has!

### Separated modules

I try to split the modules in different packages.
So far we have one for:
* the context: a yaml file that will help me in the future
* docker: all the docker tools to setup / start / check your docker
* gitlab: all you need to setup you gitlab (ssh key, ssh config, key in gitlab, test the connection)
* sdkman: helps you to install sdkman
* more to come...
