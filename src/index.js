const server = require('./server')
const config = require('./config')

const app = server(config, { logger: true })

app.listen(3000, '::', (err, address) => {
  if (err) {
    app.log.error(err)
    process.exit(1)
  }
})
