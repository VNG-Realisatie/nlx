module.exports = {
  extends: [
    'standard',
    'plugin:prettier/recommended',
    '@commonground/eslint-config/rules/reactAppGenerics',
    '@commonground/eslint-config/rules/generic',
    '@commonground/eslint-config/rules/prettier',
    '@commonground/eslint-config/rules/header',
    '@commonground/eslint-config/rules/import',
  ],

  globals: {
    fixture: 'readonly',
    test: 'readonly',
  },
}
