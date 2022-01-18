## go-places

This is probably one of the most useless API's there is, but it's not intended to be useful. 
It's mere purpose is to demonstrate how to create a workload written in `golang` in Tanzu Application Platform.


kubectl create secret generic mariadb-secret --from-literal=MARIADB_USER=dbuser --from-literal=MARIADB_PASSWORD=secretpass

kubectl apply -f .

### Create a workload

Create the workload in TAP:

```
tanzu apps workload create go-places \
     --git-repo https://github.com/pvdbleek/go-places \ 
     --git-branch main \
     --type web \
     --label app.kubernetes.io/part-of=go-places \
     --yes
```

Watch it build and deploy:

```
tanzu apps workload tail go-places
```

Once done, fetch the URL:

```
tanzu apps workload get go-places
```

Which should result in something like this:

```
# go-places: Ready
---
lastTransitionTime: "2022-01-14T20:00:35Z"
message: ""
reason: Ready
status: "True"
type: Ready

Workload pods
NAME                                STATE       AGE
go-places-build-1-build-pod         Succeeded   10m
go-places-config-writer-dhpwc-pod   Succeeded   9m12s

Workload Knative Services
NAME        READY   URL
go-places   Ready   http://go-places.default.192.168.64.6.nip.io
```

The repo also has a `catalog.yaml` which can be registered in your TAP GUI Catalogs.
When TAP GUI asks for the Repository URL, simply paste this and hit analyze:

```
https://github.com/pvdbleek/go-places/blob/main/catalog-info.yaml
```
### Using the API

The API has a the following endpoints:

| Endpoint     | Method      | Description                                |
| ------------ | ----------- | ------------------------------------------ |
| /places      | GET         | Fetches all places                         |
| /places/{id} | GET         | Fetches a specific place                   |
| /places      | POST        | Add/replace a place                        |
| /url/{id}    | GET         | Generates a Google Maps URL for this place |

### Example POST

```
curl http://go-places.default.192.168.64.6.nip.io/places \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "3","name": "Heaven's Gate","country": "China","description": "A stairway to heaven on Tianmen Mountain","latitude": 29.053743429510085,"longitude": 110.48154034958873}'
```

## Known issue(s)

Any POSTs you make are only persisted in memory, so it doesn't run nicely in knative yet. Once you make a change or post a new place, it will probably not exist in the next request because knative spun up a new instance.

Need to add a database or so to persist data.