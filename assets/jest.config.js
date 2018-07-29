module.exports = {
    verbose: true,
    testURL: 'http://localhost/',
    transform: {
        // '^.+\\.vue$': 'vue-jest',
        // '.+\\.(css|styl|less|sass|scss|png|jpg|ttf|woff|woff2)$':
        //   'jest-transform-stub',
        '^.+\\.(js|jsx)?$': 'babel-jest'
      },
    "testPathIgnorePatterns": [
        "node_modules",
        "/.yarn/",
      ],
    "moduleNameMapper": {
        "^css(.+)$": "<rootDir>/css$1",
        "src(.+)$": "<rootDir>/src$1"
    }
}
