// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../test-utils'
import Avatar from './index'

describe('Avatar', () => {
  it('should render the provided image', () => {
    const { container } = renderWithProviders(<Avatar url="my-avatar.png" />)
    const image = container.querySelector('img')
    expect(image.getAttribute('src')).toBe('my-avatar.png')
  })

  it('should render the default avatar if no image is provided', () => {
    const { container } = renderWithProviders(<Avatar />)
    const image = container.querySelector('img')
    expect(image.getAttribute('src')).toBe('default-avatar.svg')
  })
})
