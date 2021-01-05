// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders, fireEvent } from '../../../../test-utils'
import ModalHeader from './index'

test('renders title and close button', () => {
  const handleUserClose = jest.fn()

  const { container, getByText, getByTitle } = renderWithProviders(
    <ModalHeader
      handleUserClose={handleUserClose}
      title="The title"
      showCloseButton
    />,
  )

  // eslint-disable-next-line no-useless-concat
  expect(container.textContent).toBe('The title' + 'close.svg')
  expect(getByText('The title')).toBeInTheDocument()

  fireEvent.click(getByTitle('Close'))
  expect(handleUserClose).toHaveBeenCalled()
})

test('renders empty header element', () => {
  const handleUserClose = jest.fn()

  const { container } = renderWithProviders(
    <ModalHeader handleUserClose={handleUserClose} showCloseButton={false} />,
  )

  const header = container.querySelector('header')
  expect(header).toBeInTheDocument()
  expect(header.textContent).toBe('')
})
