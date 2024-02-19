# Stori technical assesment

## Requirements

- [Docker](https://www.docker.com/)
- [Git](https://git-scm.com/downloads)

## Installation

Clone the repository

```sh
git clone https://github.com/casmelad/stori.git
```
## Docker

By default, the Docker will expose port 5432 for the postgres server, so change this within the
docker-compose.yml if necessary. 

```sh
    ports:
      - 5432:5432
```

When ready, to build the images and run the container use the command as shown

```sh
docker compose up
```

once the application is running it will keep runing during 1 minute sending the email every 10 seconds

## Environment variables

You can edit the next environment variables in teh docker-compose.yaml



```sh
- SMTP_PASSWORD=value
- SMTP_USEREMAIL=value
- SMTP_HOST=value
- SMTP_PORT=value
```

There are more environment variables that are not recommended to be modified.

## Test file

You can always modify the csv file in order to test the email calculations

[CSV File](https://github.com/casmelad/stori/blob/main/txns.csv)


## License

Apache 2.0
