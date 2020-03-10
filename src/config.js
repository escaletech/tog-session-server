const sanitizePath = path =>
  path ? path.replace(/^\/(.*)\/$/, '/$1') : ''

const defaultExpiration = Number(process.env.DEFAULT_EXPIRATION) || 60 * 60 * 24 * 7

module.exports = Object.freeze({
  pathPrefix: sanitizePath(process.env.PATH_PREFIX || ''),
  redisUrl: process.env.REDIS_URL || 'redis://127.0.0.1:6379',
  isRedisCluster: process.env.REDIS_CLUSTER === 'true',
  defaultExpiration,
  maxDuration: Number(process.env.MAX_DURATION) || defaultExpiration
})
