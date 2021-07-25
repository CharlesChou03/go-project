export PATH="$(go env GOPATH)/bin:$PATH"
export MONGO_URI="mongodb://localhost:27016/shorten-url?rm.failover=1000ms:5x1&rm.monitorRefreshMS=100&rm.nbChannelsPerNode=1"