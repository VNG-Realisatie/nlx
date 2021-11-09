// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
module.exports = {
  extends: '@commonground/eslint-config-cra-standard-prettier',
  overrides: [
    {
      files: ['*.test.js', '*.test.jsx', '*.test.ts', '*.test.tsx'],
      rules: {
        'react/display-name': 0,
        '@typescript-eslint/no-empty-function': 0,
      },
    },
  ],
}
