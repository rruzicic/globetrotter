# Automated scripts

## `delete_mongo_data.py`

`data` folder is where MongoDB stores it's data. Since MongoDB is running in container the data folder on host system is bound with the data folder in 
MongoDB container so any manipulation of `./data` will result in changing `/data/db` folder inside of container. 

**WARNING**: Executing this script while the container is running will result with MongoDB container exiting with status code 14.