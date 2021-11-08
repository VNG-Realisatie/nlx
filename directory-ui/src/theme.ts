// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { defaultTheme } from '@commonground/design-system'

const tokens = {
  ...defaultTheme.tokens,

  containerWidth: '992px',

  colors: {
    colorPaletteGray200: '#EEEEEE',
    colorPaletteGray300: '#E0E0E0',
  },
}

const theme = {
  ...defaultTheme,

  tokens,

  // Table
  colorBorderTable: tokens.colorPaletteGray200,
  colorBackgroundTableHover: tokens.colorPaletteGray200,
  colorBackgroundTableSelected: tokens.colorPaletteGray300,
}

export default theme
