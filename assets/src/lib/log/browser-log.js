/**
 * Environment variables.
 *
 * @type {String}
 */
import moment from 'moment'

const NODE_ENV = typeof process !== 'undefined' && process.env && process.env.NODE_ENV
const LOG_LEVEL = localStorage.getItem('LOG_LEVEL') || (NODE_ENV === 'production' ? 'error' : 'debug')

// use localStorage.LOG_PREFIX to switch log

/**
 * Log levels.
 *
 * @type {Object}
 */

const LEVELS = {
    debug: 1,
    info: 2,
    trace: 3,
    warn: 4,
    error: 5,
}

/**
 * Log level styles.
 *
 * @type {Object}
 */

const STYLES = {
    debug: 'color: #9AA2AA',
    info: 'color: #659AD2',
    trace: 'color: blue',
    warn: 'color: #F9C749',
    error: 'color: #EC3D47',
}

const colors = [
    '#0000CC', '#0000FF', '#0033CC', '#0033FF', '#0066CC', '#0066FF', '#0099CC',
    '#0099FF', '#00CC00', '#00CC33', '#00CC66', '#00CC99', '#00CCCC', '#00CCFF',
    '#3300CC', '#3300FF', '#3333CC', '#3333FF', '#3366CC', '#3366FF', '#3399CC',
    '#3399FF', '#33CC00', '#33CC33', '#33CC66', '#33CC99', '#33CCCC', '#33CCFF',
    '#6600CC', '#6600FF', '#6633CC', '#6633FF', '#66CC00', '#66CC33', '#9900CC',
    '#9900FF', '#9933CC', '#9933FF', '#99CC00', '#99CC33', '#CC0000', '#CC0033',
    '#CC0066', '#CC0099', '#CC00CC', '#CC00FF', '#CC3300', '#CC3333', '#CC3366',
    '#CC3399', '#CC33CC', '#CC33FF', '#CC6600', '#CC6633', '#CC9900', '#CC9933',
    '#CCCC00', '#CCCC33', '#FF0000', '#FF0033', '#FF0066', '#FF0099', '#FF00CC',
    '#FF00FF', '#FF3300', '#FF3333', '#FF3366', '#FF3399', '#FF33CC', '#FF33FF',
    '#FF6600', '#FF6633', '#FF9900', '#FF9933', '#FFCC00', '#FFCC33'
]

const getTextColor = text => {
    const count = text.split('').reduce((count, item) => count += item.charCodeAt(0), 0)
    return colors[count % colors.length]
}

/**
 * Log level methods.
 *
 * @type {Object}
 */

const METHODS = {
    debug: console.debug ? 'debug' : 'log',
    info: console.info ? 'info' : 'log',
    // 自定义trace方法，更易理解使用warn
    trace: console.warn ? 'warn' : 'log',
    warn: console.warn ? 'warn' : 'log',
    error: console.error ? 'error' : 'log',
}


/**
 * Define the `Logger` class.
 *
 * @type {Logger}
 */

class Logger {

    /**
     * Constructor.
     *
     * @param {Object} options
     */

    constructor(options = {}) {
        let {
            level = LOG_LEVEL,
                prefix = '',
        } = options

        if (typeof level !== 'string') {
            level = 'none'
        }

        level = level.toLowerCase()

        if (!(level in LEVELS)) {
            level = 'none'
        }

        if (typeof prefix !== 'string') {
            prefix = String(prefix)
        }

        this.config = {
            level,
            prefix,
            threshold: level == 'none' ? Infinity : LEVELS[level],
            instanceColor: getTextColor(prefix)
        }

        for (const key in LEVELS) {
            this[key] = (...args) => this.log(key, ...args)
        }
    }

    isEnable = () => {
        const { prefix = '' } = this.config
        const log_prefix = localStorage.getItem('LOG_PREFIX') || '.*'
        const reg = new RegExp(log_prefix.replace(/\*/g, '.*?'))
        return reg.test(prefix)
    }

    /**
     * Log to the console with `level`
     *
     * @param {String} level
     */

    log = (level, ...args) => {
        if (!this.isEnable()) return

        if (typeof level !== 'string') {
            level = 'info'
        }

        level = level.toLowerCase()

        if (!(level in LEVELS)) {
            level = 'info'
        }

        const {
            threshold,
            prefix,
            instanceColor
        } = this.config
        const value = LEVELS[level]
        if (value < threshold) return

        const method = METHODS[level]
        const moduleString = prefix ? `[${prefix}]` : ''
        const allArgs = [`%c[${moment().format('HH:mm:ss.SSS')}]%c[${level}]%c${moduleString}`, 'color: green', STYLES[level], `color: ${instanceColor}`, ...args]
        console[method](...allArgs) // eslint-disable-line no-console
    }

    /**
     * Create a new logger, extending the current logger's config.
     *
     * @param {Object} options
     * @return {Logger}
     */

    clone = (options = {}) => new Logger({
            ...this.config,
            ...options,
        })

}

/**
 * Create a logger singleton with sane defaults.
 *
 * @type {Logger}
 */

const logger = new Logger()

/**
 * Export.
 *
 * @type {Logger}
 */

module.exports = exports = logger
exports.Logger = Logger
