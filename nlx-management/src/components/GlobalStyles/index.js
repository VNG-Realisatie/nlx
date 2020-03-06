import { createGlobalStyle } from 'styled-components'
import 'typeface-source-sans-pro/index.css'

export default createGlobalStyle`
  a, a:visited, a:hover, a:active {
    text-decoration: none;
  }
  a[disabled] {
    pointer-events: none;
  }
  
  html,
  body,
  #root {
    height: 100%;
  }

  html {
    background-color: #212121;
    color: ${(p) => p.theme.tokens.colors.colorText};
    font-family: 'Source Sans Pro', sans-serif;
    font-size: ${(p) => p.theme.tokens.baseFontSize};
    font-weight: ${(p) => p.theme.tokens.fontWeightRegular};
    line-height: ${(p) => p.theme.tokens.lineHeightText};
    text-rendering: optimizeLegibility;
    touch-action: manipulation;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  body {
    margin: 0;
    word-wrap: break-word;
    word-break: break-word;
  }
  
  #root {
    display: flex;
    flex-direction: column;
  }

  *, *:before, *:after {
    box-sizing: border-box;
  }
  
  a {
    text-decoration: underline;
    text-decoration-skip-ink: auto;
    color: ${(p) => p.theme.tokens.colors.colorTextLink};
    
    &:hover,
    &:active {
      color: ${(p) => p.theme.tokens.colors.colorTextLinkHover};
    }
  }
`
