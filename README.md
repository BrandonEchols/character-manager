# character-manager
This is an attempt to make a digital character sheet for role playing games.
This work is intended for private use and has no responsibility to anyone other than the authors. In other words, you get what you pay for :D

# Backend


## Running the server
To start the server on your local machine, you'll need to have Docker installed locally and the daemon running 
and then you can run the following in a unix-terminal:

`make run.local`

This will start up a docker image of postgres for the server to talk to, as well as run the `server` binary in the root directory

Once the server is running locally, apidocs can be found [here](http://localhost:8000/apidoc/) 

### Configurations
Configurations for the server can be found in the `internal/config/config.json`. 

If you want to alter the configs for your local env, copy `internal/config/config.json` -> `internal/config/config.json.dist`
This allows you to keep a non-checked in set of configs for your local development

## Developing on the server
The following are needed for developing on the server:

- Golang [Getting Started Guide](https://go.dev/doc/tutorial/getting-started)
- Apidoc ``npm install apidoc -g`` <-- Requires [NPM](https://nodesource.com/blog/the-basics-getting-started-with-npm) to be installed
- Docker [Docker Desktop Download](https://www.docker.com/products/docker-desktop/)
- Goose (if you're adding a new db migrations) [install guide](https://github.com/pressly/goose)

## Developing on the frontend
See the [frontend readme](./frontend/README.md) for details

### Endpoints
All Routes are defined in `internal/routes/routes.go`

### Database
Schema migration scripts and DDL can be found in the `database/migrations` folder

