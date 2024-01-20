# GO MicroServer

## Description

This project represents a simple backend server app implemented in GO + ENT (framework) + PostgreSQL. It is based on three microservices:
- db: the database 
- api: the main application
- pgadmin: the database development platform

## How to run the app:
At the root of the project:
`$ docker compose up --build`

## Technical details:

The database is initialized by `ent` (Framework Go), which generates the tables and the relationships between them. Its management is carried out with the help of an ORM. 

The database schema is defined in `ent/schema` folder and contains three tables for cities, contries and temperatures.

In the `server.go` file is the API server which will create the server and register the routers from the `router` folder. Here there are the routers for the three main types of pages (cities, countries and temperatures). These routers use the controllers defined in the `controller` folder, which in turn will call the corresponding services from `services` folder. All the dto's are specifided in the `dto` folder.

Varous utility features are implemented in `config` folder (where the ent configuration file is located), or in the `middleware`/`utils` folders.

The `Dockerfile` is used for building the API container, and the `docker-compose.yaml` file structure is described below:

- db -> the database container
   - container name: database
   - for data persistence I use the volume postgres:/var/lib/postgresql/data
   - username: admin
   - password: pass
   - database: db
- api -> the API
   - container name: api
   - accessible on localhost:6000
   - build from root -> where the Dockerfile is located
- pgadmin -> the database utility
   - container name: pgadmin
   - pgadmin email: admin@admin.com
   - pgadmin password: root
   - accessible on localhost:5050
networks:
- there are 2 networks: db-network and api-network