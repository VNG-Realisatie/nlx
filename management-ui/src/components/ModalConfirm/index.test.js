// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, act } from '@testing-library/react'

import { renderWithProviders } from '../../test-utils'
import { useModalConfirm } from './index'

// eslint-disable-next-line react/prop-types
const TestCase = ({ handleChoice }) => {
  const [ModalConfirm, confirm] = useModalConfirm({
    children: 'Weet je het zeker?',
  })

  const showConfirm = async () => {
    const choice = await confirm()
    handleChoice(choice)
  }

  return (
    <>
      <button onClick={showConfirm}>show confirm</button>
      <ModalConfirm />
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

  await act(async () => {
    fireEvent.click(getByText('Cancel'))
  })

  expect(handleChoice).toHaveBeenCalledWith(false)

  fireEvent.click(getByText('show confirm'))

  await act(async () => {
    fireEvent.click(getByText('Ok'))
  })

  expect(handleChoice).toHaveBeenCalledWith(true)
})
