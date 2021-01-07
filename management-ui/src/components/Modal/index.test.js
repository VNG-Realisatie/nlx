// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { fireEvent } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import Modal, { verticalAlignToCssValues } from './index'

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

test.concurrent.each([
  /* eslint-disable prettier/prettier */
  [
    {},
    { alignItems: 'center', transform: '' },
  ],
  [
    { from: 'top' },
    { alignItems: 'flex-start', transform: '' },
  ],
  [
    { offset: '-50%' },
    { alignItems: 'center', transform: 'translateY(-50%)' },
  ],
  [
    { from: 'bottom', offset: '100px' },
    { alignItems: 'flex-end', transform: 'translateY(-100px)' },
  ],
  [
    { from: 'bottom', offset: '-50%' },
    { alignItems: 'flex-end', transform: 'translateY(50%)' },
  ],
  /* eslint-enable prettier/prettier */
])('vertical alignment object is created as expected', (config, expected) => {
  expect(verticalAlignToCssValues(config)).toEqual(expected)
})
