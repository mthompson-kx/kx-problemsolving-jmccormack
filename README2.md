# COMMANDS:

From the kxSystems directory

A user can
    run `docker-compose up`

Or use the Makefiles commands

run `make clean`
    runs `docker-compose down` and removes images

run `make run`
    runs `docker-compose up`

run `make main`
    removes old go binaries and builds new ones

# API SERVER:

Api Server is listening on http://localhost:8000
and has endpoints
        http://localhost:8000/data
        http://localhost:8000/status


/data

    The Api Server will attempt to download a json file from one of 3 file servers
    They are selected in a round robin proccess
    If a file server fails to respond, the file will be requested from the next server in the round
    If all servers are down an `internal error message` is displayed. This is deliberatly obtuse

/status

    The Api Server will return the health status of all 3 file servers in json format

/

    Any other route will return a `PAGE_NOT_FOUND` message

# FILE SERVER

The File Servers are listening on http://localhost:8000
and has endpoints
        http://localhost:8000/data.json
        http://localhost:8000/health

/data.json

    The file server will server the static data.json file


/health

    The file server will return `ok`


The file servers are not accesible to the general public