const fastify = require('fastify')
const TogClient = require('tog-node')

const { parseOptions } = require('./options')

function server ({ pathPrefix, redisUrl, defaultExpiration }, fastifyConfig = {}) {
  const client = new TogClient(redisUrl)

  return fastify(fastifyConfig)
    .get(pathPrefix + '/:ns/:sid', (request, reply) => {
      const options = parseOptions(request.query)
      return client.session(request.params.ns, request.params.sid, defaultExpiration, options)
        .then(session => reply.status(200).send(session))
    })
}

module.exports = server
