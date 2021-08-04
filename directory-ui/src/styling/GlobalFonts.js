// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import { createGlobalStyle } from 'styled-components'

export default createGlobalStyle`
  @font-face {
    font-family: 'Source Sans Pro';
    font-style: normal;
    font-display: swap;
    font-weight: 400;
    src:
      local('Source Sans Pro Light'),
      local('SourceSansPro-Light'), url('/fonts/source-sans-pro-latin-400-normal.woff2') format('woff2'), url('/fonts/source-sans-pro-latin-400-normal.woff') format('woff');
  }

  @font-face {
    font-family: 'Source Sans Pro';
    font-style: normal;
    font-display: swap;
    font-weight: 600;
    src:
      local('Source Sans Pro Light'),
      local('SourceSansPro-Light'), url('/fonts/source-sans-pro-latin-600-normal.woff2') format('woff2'), url('/fonts/source-sans-pro-latin-600-normal.woff') format('woff');
  }

  @font-face {
    font-family: 'Source Sans Pro';
    font-style: normal;
    font-display: swap;
    font-weight: 700;
    src:
      local('Source Sans Pro Light'),
      local('SourceSansPro-Light'), url('/fonts/source-sans-pro-latin-700-normal.woff2') format('woff2'), url('/fonts/source-sans-pro-latin-700-normal.woff') format('woff');
  }

  html {
    font-family: 'Source Sans Pro', sans-serif;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }
`
