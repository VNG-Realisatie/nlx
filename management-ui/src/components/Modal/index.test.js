// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { fireEvent } from '@testing-library/react'

import { renderWithProviders } from '../../test-utils'
import Modal from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  jest.runOnlyPendingTimers()
  jest.useRealTimers()
})

test('renders a modal, closes with Escape', async () => {
  const ParentElement = () => {
    const [showModal, setShowModal] = useState(true)

    return (
      <div>
        <Modal
          isVisible={showModal}
          handleClose={() => {
            setShowModal(false)
          }}
          title="The title"
        >
          <p>Modal content</p>
        </Modal>
      </div>
    )
  }

  const { getByText } = renderWithProviders(<ParentElement />)
  const content = getByText('Modal content')

  expect(getByText('The title')).toBeInTheDocument()
  expect(content).toBeInTheDocument()

  fireEvent(document, new KeyboardEvent('keydown', { key: 'Escape' }))
  jest.runAllTimers()

  expect(content).not.toBeInTheDocument()
})
