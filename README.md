# cookbook-api
This is the API for the Cookbook project
## Setting up Docker container locally
1. Install `docker`, `virtualbox`, and `docker-compose`
2. Run `docker-machine start default` to start Docker machine
3. Run `eval $(docker-machine env default)`
4. Copy `docker-compose.yml.default` to `docker-compose-yml`
5. Fill in database information in new docker-compose file
6. Run `docker-compose up -d --build` to start application and database