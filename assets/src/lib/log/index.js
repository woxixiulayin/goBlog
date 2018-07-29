const logger = require('./browser-log')

const log = logger.info
Object.assign(log, logger)

module.exports = log