const fastify = require('fastify')
const { SessionClient } = require('tog-node')

const { parseOptions } = require('./options')

function server ({ pathPrefix, redisUrl, defaultExpiration }, fastifyConfig = {}) {
  const client = new SessionClient(redisUrl)

  return fastify(fastifyConfig)
    .get(pathPrefix + '/:ns/:sid', (request, reply) => {
      const options = parseOptions(request.query)
      return client.session(request.params.ns, request.params.sid, {
        flags: options.flags,
        duration: defaultExpiration
      })
        .then(session => reply.status(200).send(session))
    })
}

module.exports = server
