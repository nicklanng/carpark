# carpark
An HTTP API for paying for parking.

[![Build Status](https://travis-ci.org/nicklanng/carpark.svg?branch=master)](https://travis-ci.org/nicklanng/carpark)

## Setup
This application is written in Go and tested with Ruby and RSpec. The tests run against docker containers.

- Golang 1.10.x https://golang.org/
- Ruby 2.4.x https://www.ruby-lang.org/
- Bundler gem http://bundler.io/
- Docker https://www.docker.com/
- Protobuf 3.5.1 https://github.com/google/protobuf

Run this command to install all required Ruby gems and go tools.
```bash
$ bundle install && rake install
```

Available Rake tasks:
```bash
rake install     # Go get required tools
rake clean       # Remove any temporary products
rake codegen     # Generate protobuf and stringer go code
rake unittest    # Run Golang unit tests
rake build       # Build a local docker image
rake rubycodegen # Generate protobuf ruby code
rake spec        # Run RSpec code examples
```
