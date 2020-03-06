const tokens = {
  baseFontSize: '16px',

  lineHeightText: '150%',
  lineHeightHeading: '125%',

  fontWeightRegular: '500',
  fontWeightSemiBold: '600',
  fontWeightBold: '700',

  fontSizeSmall: '0.875rem',
  fontSizeMedium: '1rem',
  fontSizeLarge: '1.125rem',
  fontSizeXLarge: '1.5rem',
  fontSizeXXLarge: '2rem',

  colors: {},
}

// Derived colors

// Text
tokens.colors.colorText = '#ffffff'
tokens.colors.colorTextLink = '#1694C8'
tokens.colors.colorTextLinkHover = '#1694C8'

const theme = {
  tokens: tokens,
}

export default theme
