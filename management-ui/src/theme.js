// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { darkTheme } from '@commonground/design-system'

const tokens = {
  ...darkTheme.tokens,
}

const theme = {
  ...darkTheme,

  tokens,

  // Dropdown
  colorBackgroundDropdown: tokens.colorPaletteGray900,
  colorBackgroundDropdownHover: '#515151',
  colorBackgroundDropdownActive: tokens.colorPaletteGray600,
  colorBorderDropdownFocus: '#1EA1D5',

  // Table
  colorBorderTable: tokens.colorPaletteGray800,
  colorBackgroundTableHover: 'rgba(255, 255, 255, 0.1)',
  colorBorderTableFocus: tokens.colorFocus,
}

export default theme
