# encryptedstore

### build and run
Ensure that:
- go is installed
- the repository is placed at `$GOPATH/src/glynternet/encryptedstore`

Then run the following:
```
go build ./cmd/encrypted-store && ./encrypted-store
```

### Usage:
With the server running locally, the following script would store a payload and then retrieve it again.
```bash
#!/bin/bash
# store request paramaters
id=foo
payload=bar
storeRequestBody="{\"id\":\"$id\",\"payload\":\"$payload\"}"
storeUrl="localhost:8080/store"

echo "id: $id"
echo "payload: $payload"
echo "request body: $storeRequestBody"
echo "store url: $storeUrl"

# the key is the body of the response
key="$(curl -X POST --data "$storeRequestBody" "$storeUrl")"

# use the key to retrieve the original payload
retrieveRequestBody="{\"id\":\"$id\",\"key\":\"$key\"}"
retrieveUrl="localhost:8080/retrieve"

echo "encryption key: $key"
echo "request body: $retrieveRequestBody"
echo "retrieve url: $retrieveUrl"

retrievedPayload="$(curl -X POST --data "$retrieveRequestBody" "$retrieveUrl")"
echo "retrieved payload: $retrievedPayload"
```