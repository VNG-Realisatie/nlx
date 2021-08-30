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
  gradientBlue: 'linear-gradient(135deg, #295372, #163145)',
  colorDarkBlue: '#30709d',
  colorAverageBlue: '#20425c',
  colorAlternateSection: '#f1f1f1',
  listIconSize: '2.5rem',

  colorCollapsibleBorder: tokens.colors.colorPaletteGray300,

  breakpoints: Object.values(breakpoints)
    .splice(1)
    .map((bp) => `${bp}px`),

  // Shared
  colorFocus: '#1EA1D5',

  // Dropdown
  colorBackgroundDropdown: tokens.colorPaletteGray900,
  colorBackgroundDropdownHover: '#515151',
  colorBackgroundDropdownActive: tokens.colorPaletteGray600,

  // Table
  colorBorderTable: tokens.colorPaletteGray200,
  colorBackgroundTableHover: tokens.colorPaletteGray200,
  colorBackgroundTableSelected: tokens.colorPaletteGray300,
  colorBorderTableFocus: tokens.colorFocus,
}

theme.colorBackgroundDropdown = theme.colorBackgroundInput
theme.colorBackgroundDropdownHover = tokens.colorPaletteGray200
theme.colorBackgroundDropdownActive = tokens.colorPaletteGray400
theme.colorTextDropdownActive = tokens.colorPaletteGray900
theme.colorBackgroundDropdownSelected = tokens.colorPaletteGray300

export default theme
