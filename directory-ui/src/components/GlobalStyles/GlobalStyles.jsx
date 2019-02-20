import { createGlobalStyle } from 'styled-components'

const generateFontFaceDefinition = fontDefinition => `
  @font-face {
      font-family: 'Source Sans Pro'
      font-weight: ${fontDefinition.weight}
      font-style: ${fontDefinition.style}
      font-stretch: normal
      src:
        url('/public/fonts/source-sans-pro/SourceSansPro-${fontDefinition.fileName}.woff2') format('woff2'),
        url('/public/fonts/source-sans-pro/SourceSansPro-${fontDefinition.fileName}.woff') format('woff')
  }
`

const generateFontFaceDefinitions = fontDefinitions =>
  fontDefinitions.map(fontDefinition => generateFontFaceDefinition(fontDefinition))

const fontDefinitions = [
  { weight: 400, style: 'normal', fileName: 'Regular' },
  { weight: 600, style: 'normal', fileName: 'Semibold' },
  { weight: 700, style: 'normal', fileName: 'Bold' },
]

export default createGlobalStyle`
  ${generateFontFaceDefinitions(fontDefinitions)}

  html,
  body {
    background: #F7F9FC;
    font-family: 'Source Sans Pro', sans-serif;
    font-size: 14px;
  }

  body {
    margin: 0;
  }
`
