import { createGlobalStyle } from 'styled-components'

const generateFontFaceDefinition = fontDefinition => `
  @font-face {
      font-family: 'Source Sans Pro'
      font-weight: ${fontDefinition.weight}
      font-style: ${fontDefinition.style}
      font-stretch: normal
      src: url('/public/fonts/source-sans-pro/SourceSansPro-${fontDefinition.fileName}.otf.woff2') format('woff2')
  }
`

const generateFontFaceDefinitions = fontDefinitions =>
  fontDefinitions.map(fontDefinition => generateFontFaceDefinition(fontDefinition))

const fontDefinitions = [
  { weight: 200, style: 'normal', fileName: 'ExtraLight' },
  { weight: 200, style: 'italic', fileName: 'ExtraLightIt' },
  { weight: 300, style: 'normal', fileName: 'Light' },
  { weight: 300, style: 'italic', fileName: 'LightIt' },
  { weight: 400, style: 'normal', fileName: 'Regular' },
  { weight: 400, style: 'italic', fileName: 'It' },
  { weight: 600, style: 'normal', fileName: 'Semibold' },
  { weight: 600, style: 'italic', fileName: 'SemiboldIt' },
  { weight: 700, style: 'normal', fileName: 'Bold' },
  { weight: 700, style: 'italic', fileName: 'BoldIt' },
  { weight: 900, style: 'normal', fileName: 'Black' },
  { weight: 900, style: 'italic', fileName: 'BlackIt' },
]

export default createGlobalStyle`
  ${generateFontFaceDefinitions(fontDefinitions)}

  html,
  body {
    background: #F7F9FC;
    font-family: 'Source Sans Pro', sans-serif;
    font-size: 14px;
  }
`
