# Tog Session Server

[![CircleCI](https://circleci.com/gh/escaletech/tog-session-server.svg?style=svg)](https://circleci.com/gh/escaletech/tog-session-server)

Server application that provides an endpoint for
fetching [Tog](https://github.com/escaletech/tog) sessions.

## Usage

```sh
$ docker run -d -p 3000:3000 \
  --env 'REDIS_URL=redis://your-redis:6379' \
  escaletech/tog-session-server
$ curl localhost:3000/<NAMESPACE>/<SESSION-ID>
```

### Configuration variables

* `PATH_PREFIX` - Prefix to the endpoint (**optional**, e.g. `/_sessions`)
* `REDIS_URL` - URL for the Redis server used by Tog (**required**, e.g. `redis://my-redis-server.com`)
* `REDIS_CLUSTER` - Set to `true` if Redis URL is a cluster (**optional**, default: `false`)
* `DEFAULT_EXPIRATION` - Expiration (**optional**, default: `604800` - one week)
