# Goland
Goland is a simple web server built in golang using gin web framework.

#### Frameworks used

|Name  | Use  |
|:-----| :--- |
|[gin](https://github.com/gin-gonic/gin)| Web frame work|
|[logrus](https://github.com/sirupsen/logrus)|Structured logger |
|[gomock](https://github.com/golang/mock)|Mocking framework for testing|
|[dlv](https://github.com/go-delve/delve/tree/master/cmd/dlv)|Debugger|
|[swagger](https://github.com/go-swagger/go-swagger)|REST API representation tooling|
-----------------
#### Features
Following are the list of features supported in Goland.
```bash
compile                        Build the app
fmt                            Run the code formatter
generate-docs                  Generates the static files to be embedded into the application + swagger.json
generate-mocks                 Generate mocks to be used only for unit testing
lint                           Run the code linter
run                            Build and start app locally (outside docker)
setup                          Setup necessary dependencies and folder structure
test                           Run tests
```

#### Installation
To setup Goland web server use the following command, this uses the Makefile with predefined 
configuration.

```bash
make setup build-deps build && echo $?
```

#### Run the server
To start the server use the following, by default server runs in `8000`.

```bash
make run
```
One can customize the port us follows
```bash
PORT=8001 make run
```

