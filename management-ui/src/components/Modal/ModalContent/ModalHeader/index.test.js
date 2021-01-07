// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders, fireEvent } from '../../../../test-utils'
import ModalHeader from './index'

test('renders title and close button', () => {
  const handleClose = jest.fn()

  const { container, getByText, getByTitle } = renderWithProviders(
    <ModalHeader onClose={handleClose} title="The title" showCloseButton />,
  )

  // eslint-disable-next-line no-useless-concat
  expect(container.textContent).toBe('The title' + 'close.svg')
  expect(getByText('The title')).toBeInTheDocument()

  fireEvent.click(getByTitle('Close'))
  expect(handleClose).toHaveBeenCalled()
})

test('renders empty header element when expected', () => {
  const handleClose = jest.fn()

  const { container } = renderWithProviders(
    <ModalHeader onClose={handleClose} showCloseButton={false} />,
  )

  const header = container.querySelector('header')
  expect(header).toBeInTheDocument()
  expect(header.textContent).toBe('')
})
