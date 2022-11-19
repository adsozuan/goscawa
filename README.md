# goscawa

A simple go web api example applying clean architecure for working out Go 

* As simple as possible
* Gin as dependency
* No surprise
* Extensible


Largely inspired by Mat Ryer videos on how he writes web API and Clean Architecture principles practices gathered along my career.

## Specification and domain

For demonstration purpose, the web API allows to found the nearest waste disposal near your home.

> Note: for France only ;).

For this  will use [ADEME directory of waste disposal center](https://data.ademe.fr/applications/sinoe-(r)-annuaire-des-decheteries-dma-infos-localisations?lat=39.88034212135295&lng=-2.5662082591675244&zoom=2.6090413222127125) web site service. ADEME is the French Ecological Transition Agency.

A OpenAPI definition is provided here [api_docs](https://data.ademe.fr/data-fair/api/v1/datasets/sinoe-(r)-annuaire-des-decheteries-dma/api-docs.json).

- https://data.ademe.fr/data-fair/api/v1/datasets/sinoe-(r)-annuaire-des-decheteries-dma/lines


### Use cases

* A client searching for the **nearest waste disposal center** send a request to `/waste-disposal/api/nearest` with its geo coordinates. The web api, first request to ADEME gateway necessary informations to compute the nearest center.

### Domain

* WasteDisposalCenter
* GeoLocation
* FindNearest

### External System

* An adapter to ADEME API


# Build

```shell
make build
```
Will create an executable named `server` in bin/ folder.

With Docker:

```shell
./scripts/build_image.sh
```
An image named `goscawa` will be build.

# Run

```shell
make run
```

To test it:
```shell
curl localhost:8080?name=Dude
```

# Deploy on K8S

Tested with minikube.

Make you docker client point to run in minikube environment
with:

```shell
eval $(minikube docker-env)
```

Build your image:

```shell
./scripts/build_image.sh
```

Deploy it:

```shell
minikube kubectl -- apply -f build/package/deployment.yaml
```



