const sanitizePath = path =>
  path ? path.replace(/^\/(.*)\/$/, '/$1') : ''

module.exports = Object.freeze({
  pathPrefix: sanitizePath(process.env.PATH_PREFIX || ''),
  redisUrl: process.env.REDIS_URL || 'redis://127.0.0.1:6379',
  defaultExpiration: Number(process.env.DEFAULT_EXPIRATION) || 60 * 60 * 24 * 7
})
