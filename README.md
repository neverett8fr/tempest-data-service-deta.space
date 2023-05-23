# tempest-user-service  
this service is responsible for handling data storage and retrieval operations.  
  
# Features  
- data upload 
- data retrieval   

  
# How to run  
this application contains a `Dockerfile` - this allows you run build and run the service using Docker console commands   
## Build  
```bash
docker build -t .
 ```
   
 ## Run  
 ```bash
docker run -p 8080:8080 -v . -e ENV_VARIABLE=value .
 ```
   
 ## Stop the container  
 ```bash
 docker stop container-name
 ```
## Notes  

this is the data implementation using the "deta.space" noSQL - they're free!  
this might be temporary  
  
- <https://database.deta.sh/v1/a0qc5fhriwh/tempest-data-storage/items>  

```json
{
    "key": "uuid",
    "username": "user1",
    "metadata": {
        "extension": "txt",
        "name": "FileName",
        "size": 64
    },
    "data": "encodedData"
}
```  

```json
{
    "item": {
        "key": "newUUID3",
        "username": "username2",
        "metadata":{
            "extension": "txt",
            "name":"file1",
            "size":64
        },
        "data": "encodedData"
    }
}
```
