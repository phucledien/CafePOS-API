# CafePOS API

**CafePOS API** is a golang implement(using go-kit, gorm) of CafePOS API

Submitted by: **Phucledien**

Time spent: **0** hours spent in total

## Prerequisite

* Must Read about [git flows](https://nvie.com/posts/a-successful-git-branching-model)
* Should read about Postgres database, [go-kit](https://gokit.io/examples/stringsvc.html), [GORM](http://doc.gorm.io/)
* Already Setup: `git`, `golang`, [docker](https://docs.docker.com/install), [docker-compose](https://docs.docker.com/compose/install/#install-compose)

## How to start

* `make local-env`: To create local DB(port: 5432), test DB(port:5439), adminer(tools to view database, port:8080)
* `make clean-local-env`: Turn off `local-env` (be careful it will also clear DB)
* `make dev`: To start server (default port is 3000)
* `make test`: To run test (both integration and unit test)

## User Stories

### Required:
* NOTE:  Each main point should follow `git flows` (each step should have a feature branch,...)
## Notes

Notes about current git.

## License

    Copyright [2018] [phucledien]

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
