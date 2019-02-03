this is goa gRPC example for find bugs or check issues.

to start server.

```
go get -u goa.design/goa/...
./start.sh 
```


to check via grpcc client

```
npm install -g grpcc
./client.sh
```

result

```
Error:  { Error: 2 UNKNOWN: Stream removed
    at Object.exports.createStatusError (/Users/h_ayabe/.nvm/versions/node/v7.10.0/lib/node_modules/grpcc/node_modules/grpc/src/common.js:91:15)
    at Object.onReceiveStatus (/Users/h_ayabe/.nvm/versions/node/v7.10.0/lib/node_modules/grpcc/node_modules/grpc/src/client_interceptors.js:1204:28)
    at InterceptingListener._callNext (/Users/h_ayabe/.nvm/versions/node/v7.10.0/lib/node_modules/grpcc/node_modules/grpc/src/client_interceptors.js:568:42)
    at InterceptingListener.onReceiveStatus (/Users/h_ayabe/.nvm/versions/node/v7.10.0/lib/node_modules/grpcc/node_modules/grpc/src/client_interceptors.js:618:8)
    at callback (/Users/h_ayabe/.nvm/versions/node/v7.10.0/lib/node_modules/grpcc/node_modules/grpc/src/client_interceptors.js:845:24) code: 2, metadata: {}, details: 'Stream removed' }
```
