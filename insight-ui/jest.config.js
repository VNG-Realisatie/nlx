module.exports = {
    setupTestFrameworkScriptFile: '<rootDir>/src/utils/testing/setupEnzyme.js',
    moduleNameMapper: {
        '\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$':
            '<rootDir>/src/utils/testing/fileMock.js',
        '\\.(css|less)$': 'identity-obj-proxy',
    },
}
