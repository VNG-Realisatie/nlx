// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent } from '@testing-library/react'

import { renderWithProviders } from '../../../../test-utils'
import DirectoryDetailView from './index'

const mockRequestAccess = jest.fn()

test('should have a button to request access', () => {
  const { getByText } = renderWithProviders(
    <DirectoryDetailView
      onRequestAccess={mockRequestAccess}
      isAccessRequested={false}
    />,
  )

  const button = getByText('Request Access')
  expect(button).toBeInTheDocument()

  fireEvent.click(button)
  expect(mockRequestAccess).toHaveBeenCalled()
})
