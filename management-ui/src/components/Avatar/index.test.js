import React from 'react'
import { render } from '@testing-library/react'
import { ThemeProvider } from 'styled-components/macro'

import theme from '../../theme'
import Avatar from './index'

describe('Avatar', () => {
  it('should render the provided image', () => {
    const { container } = render(
      <ThemeProvider theme={theme}>
        <Avatar url="my-avatar.png" />
      </ThemeProvider>,
    )
    const image = container.querySelector('img')
    expect(image.getAttribute('src')).toBe('my-avatar.png')
  })

  it('should render the default avatar if no image is provided', () => {
    const { container } = render(
      <ThemeProvider theme={theme}>
        <Avatar />
      </ThemeProvider>,
    )
    const image = container.querySelector('img')
    expect(image.getAttribute('src')).toBe('default-avatar.svg')
  })
})
