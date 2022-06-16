# Studing Algorithms and Data Structures in some languages

## Languages choosing

- [Javascript (Node)](https://nodejs.org/)
- [PHP](https://www.php.net/)
- [Python](https://www.python.org/)
- [Golang](https://go.dev/)
- [Elixir](https://elixir-lang.org/)

## Pre-requirement to run the code

- [docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)

## Running The code

In order to maintain the things simple, I configurated a node project to run script with `npm` command.

All the code run in docker containers.

`docker-compose up --build`

**Example Node**
```bash
$ docker-compose up node-code
```
**Example PHP**
```bash
$ docker-compose up php-code
```
**Example Python**
```bash
$ docker-compose up python-code
```
**Example Golang**
```bash
$ docker-compose up golang-code
```
**Example Elixir**
```bash
$ docker-compose up elixir-code
```
