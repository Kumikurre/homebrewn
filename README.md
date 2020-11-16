# homebrewn
IoT course homebrewery monitoring system

## Instructions on running
To spin up the docker container running backend, web server and the frontend:
`docker-compose build && docker-compose up`
To test backend, run client (posts devices to database and gets them)
`cd client`
`python3 client.py`

## Components
### Backend
The backend consists of a nginx-web server which runs a GO-backend.
### Frontend
The nginx web server also serves a mithril-based frontend.
### Gateway
### Client
Just a python script for now
