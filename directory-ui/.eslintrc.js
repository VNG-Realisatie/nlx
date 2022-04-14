// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//

const isDev =
  process &&
  process.env &&
  process.env.NODE_ENV &&
  process.env.NODE_ENV === 'development'

module.exports = {
  parserOptions: {
    ecmaFeatures: {
      jsx: true,
    },
  },

  plugins: ['header', 'jest', 'import', 'prettier'],

  extends: [
    'react-app',
    'plugin:react/recommended',
    'plugin:prettier/recommended',
    'plugin:security/recommended',
  ],
  rules: {
    // generic
    'no-console': isDev
      ? ['warn', { allow: ['error', 'warn'] }]
      : ['error', { allow: ['error', 'warn'] }],
    'no-unused-vars': isDev ? 'warn' : 'error',

    // prettier
    'prettier/prettier': [
      'error',
      {
        tabWidth: 2,
        useTabs: false,
        semi: false,
        singleQuote: true,
        trailingComma: 'all',
        arrowParens: 'always',
      },
    ],

    // header
    'header/header': [
      2,
      'line',
      [
        {
          pattern: ' Copyright © VNG Realisatie \\d{4}$',
          template: ' Copyright © VNG Realisatie 2022',
        },
        ' Licensed under the EUPL',
        '',
      ],
    ],

    // react
    'react/no-unsafe': 'warn',
    'react/forbid-prop-types': ['error', { forbid: ['any'] }],
    'react/jsx-handler-names': 'error',
    'react/jsx-pascal-case': ['error', { allowAllCaps: true }],
    'react/sort-comp': 'error',

    // import
    'import/order': 'error',
    'import/first': 'error',
    'import/no-amd': 'error',
    'import/no-webpack-loader-syntax': 'error',

    // jest
    'jest/consistent-test-it': [
      'error',
      {
        fn: 'test',
        withinDescribe: 'it',
      },
    ],
    'jest/expect-expect': [
      'error',
      {
        assertFunctionNames: ['expect'],
      },
    ],
    'jest/no-jasmine-globals': 'error',
    'jest/no-done-callback': 'error',
    'jest/prefer-to-contain': 'error',
    'jest/prefer-to-have-length': 'error',
    'jest/valid-describe-callback': 'error',
    'jest/valid-expect-in-promise': 'error',

    // security
    'security/detect-object-injection': 'error',
  },
  overrides: [
    {
      files: ['*.test.js', ',*.test.jsx'],
      rules: {
        'react/display-name': 0,
      },
    },
  ],
}
