// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../../../test-utils'
import OrderDetailView from './index'

const order = {
  reference: 'my-reference',
  delegatee: 'delegatee',
  validFrom: new Date(),
  validUntil: new Date(),
}

test('display order details', () => {
  const { queryByText } = renderWithProviders(
    <OrderDetailView order={order} revokeHandler={() => {}} />,
  )

  expect(queryByText('my-reference')).toBeInTheDocument()

  // it('should call the removeHandler on remove', async () => {
  //   const handleRemove = jest.fn()
  //   const { getByTitle, getByRole } = renderWithProviders(
  //     <Router>
  //       <OrderDetailView service={order} removeHandler={handleRemove} />
  //     </Router>,
  //   )
  //
  //   fireEvent.click(getByTitle('Remove service'))
  //
  //   const confirmModal = getByRole('dialog')
  //   const okButton = within(confirmModal).getByText('Remove')
  //
  //   fireEvent.click(okButton)
  //   await waitFor(() => expect(handleRemove).toHaveBeenCalled())
  // })
})
