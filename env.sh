export PATH="$(go env GOPATH)/bin:$PATH"
export MONGO_URI="mongodb://localhost:27016/shorten-url?rm.failover=1000ms:5x1&rm.monitorRefreshMS=100&rm.nbChannelsPerNode=1"
export SHORTENING_URL_LENGTH=6
export URL_HOST="http://localhost:9999/"
export DEFAULT_EXPIRED_PERIOD=157680000000