# Cheatsheet

## Start docker containers

To start all docker containers execute `docker compose up`, if this is your first time starting the containers you should build them first, you can do this by executing 
`docker compose build` and then start them with `docker compose up` or alternatively you can combine these two commands into one: `docker compose up --build`.

To stop the containers and remove them run: `docker compose down`

## List all data in docker mongo image 

`docker exec mongo bash -c "mongosh <db_name> --quiet --eval 'db.<collection_name>.find();'"`

### Queries for existing databases(expand this list as needed)
- ` docker exec mongo bash -c "mongosh flights --quiet --eval 'db.users.find();'" `

# Automated scripts

## `delete_mongo_data.py`

`data` folder is where MongoDB stores it's data. Since MongoDB is running in container the data folder on host system is bound with the data folder in 
MongoDB container so any manipulation of `./data` will result in changing `/data/db` folder inside of container. 

**WARNING**: Executing this script while the container is running will result with MongoDB container exiting with status code 14.

