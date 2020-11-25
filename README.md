# homebrewn
IoT course homebrewery monitoring system

## Instructions on running
To spin up the docker container running backend, web server and the frontend:

`docker-compose up --build`

To spin up dev env in docker, run:

`docker-compose -f docker-compose.yml -f docker-compose-dev.yml up --build`

To test backend, run a simple test client (posts devices to database and gets them)

`cd config`

`python3 client.py`

## Components
### Backend
The backend consists of a nginx-web server which runs a GO-backend and mongo database.
### Frontend
The nginx web server also serves a mithril-based frontend.
### Client
