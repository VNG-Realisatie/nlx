// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import { clickConfirmButtonAndAssert } from './testUtils'
import { useConfirmationModal } from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

afterEach(() => {
  jest.runOnlyPendingTimers()
  jest.useRealTimers()
})

// eslint-disable-next-line react/prop-types
const TestCase = ({ handleChoice }) => {
  const [ConfirmationModal, confirmModal] = useConfirmationModal({
    children: 'Weet je het zeker?',
  })

  const showConfirm = async () => {
    const choice = await confirmModal()
    handleChoice(choice)
  }

  return (
    <>
      <button onClick={showConfirm}>show confirm</button>
      <ConfirmationModal />
    </>
  )
}

test('Interact with confirm window', async () => {
  const handleChoice = jest.fn()
  const { getByText } = renderWithProviders(
    <TestCase handleChoice={handleChoice} />,
  )

  fireEvent.click(getByText('show confirm'))

  expect(getByText('Weet je het zeker?')).toBeInTheDocument()

  await clickConfirmButtonAndAssert(getByText('Cancel'), () =>
    expect(handleChoice).toHaveBeenCalledWith(false),
  )

  fireEvent.click(getByText('show confirm'))

  await clickConfirmButtonAndAssert(getByText('Ok'), () =>
    expect(handleChoice).toHaveBeenCalledWith(true),
  )

  expect(handleChoice).toHaveBeenCalledTimes(2)
})
