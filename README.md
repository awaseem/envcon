<pre>
   _____  ___   ___________  _  __
  / __/ |/ / | / / ___/ __ \/ |/ /
 / _//    /| |/ / /__/ /_/ /    / 
/___/_/|_/ |___/\___/\____/_/|_/
</pre>

# ENVCON
[![Go Report Card](https://goreportcard.com/badge/github.com/awaseem/envcon)](https://goreportcard.com/report/github.com/awaseem/envcon)


Separate environment variable dependencies without polluting your local environment workspace.
Idea is to keep each applications environment variables as separate containers, think of `virtualenv` for environment variables.

# Install

Currently supported on OSX only. Download the binary [here](https://github.com/awaseem/envcon/releases)

## Commands

### create
create a new container with environment variables, you can also encrypt the contents.

![](./static/envcon_create.gif)

### source
create a new process with all the variables sourced.

![](./static/envcon_source.gif)

### update
update environment variable keys

![](./static/envcon_update.gif)

### delete 
delete a container

![](./static/envcon_delete.gif)

## TODO

- Tests for Commands
- Test support for linux
- Add support for Windows