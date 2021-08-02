// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { defaultTheme } from '@commonground/design-system'

const breakpoints = defaultTheme.breakpoints

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
  homeGradient: 'linear-gradient(90deg, #d6eef9 0%, #b3d0e1 100%)',
  homeGradientMobile: 'linear-gradient(90deg, #d6eef9 0%, #cbe6f3 100%)',
  gradientBlue: 'linear-gradient(135deg, #295372, #163145)',
  colorAverageBlue: '#20425c',
  colorAlternateSection: '#f1f1f1',
  listIconSize: '2.5rem',

  colorCollapsibleBorder: tokens.colors.colorPaletteGray300,

  breakpoints: Object.values(breakpoints)
    .splice(1)
    .map((bp) => `${bp}px`),
}

export default theme
