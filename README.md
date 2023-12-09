# solax_exporter

A simple server that queries the Solax Cloud API for inverter metrics and exports them via HTTP for prometheus scraping.

## Usage

To run the server locally, simply run:

```bash
go get .
go build -o solax_exporter .
./solax_exporter --sn <your_sn_here> --token-id <your_token_id_here>
```

For an overview of the available command line arguments, run `./solax_exporter -help`

Alternatively, this repository can be used as a docker image:

```bash
docker build . -t solax_exporter
docker run -p 9100 solax_exporter --sn <your_sn_here> --token-id <your_token_id_here>
```

A prebuild image is provided as `dubresc/solax_exporter` through [dockerhub](https://hub.docker.com/repository/docker/dubresc/solax_exporter).

Sample `docker-compose` configuration:

```docker-compose
version: '3'
services:
  solax_exporter:
    image: dubresc/solax_exporter
    ports:
      - '9100:9100'
    command: --sn <your_sn_here> --token-id <your_token_id_here>
```

## Required arguments

To query the solax cloud API, you require your inverter's SN and your API token-ID. For the former, check the [official documentation](https://www.solaxcloud.com/blue/user_api/SolaxCloud_User_Monitoring_API_V6.1.pdf) on how to find this information.

The token ID can be obtained from the [Solax Cloud web interface](https://www.solaxcloud.com). Log in and select `Service` -> `API` to generate (or look up) your token ID. This should be valid for all the inverters in your account.

## Issues

This repository has only been tested with a single `X1-Hybrid-G4` inverter. The information/format/units may differ for other inverters.

## External Documentation

* [Solax API](https://www.solaxcloud.com/blue/user_api/SolaxCloud_User_Monitoring_API_V6.1.pdf)