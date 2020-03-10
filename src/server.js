const fastify = require('fastify')
const { SessionClient } = require('tog-node')

const { parseOptions } = require('./options')

function server ({ pathPrefix, redisUrl, isRedisCluster, defaultExpiration }, fastifyConfig = {}) {
  const client = new SessionClient(redisUrl, { cluster: isRedisCluster })

  return fastify(fastifyConfig)
    .get(pathPrefix + '/:ns/:sid', (request, reply) => {
      const options = parseOptions(request.query)
      return client.session(request.params.ns, request.params.sid, {
        flags: options.flags,
        duration: isNaN(options.duration)
          ? defaultExpiration
          : options.duration
      })
        .then(session => reply.status(200).send(session))
    })
}

module.exports = server
