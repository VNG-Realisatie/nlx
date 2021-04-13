// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import userEvent from '@testing-library/user-event'
import { waitFor, fireEvent } from '@testing-library/react'
import selectEvent from 'react-select-event'
import { renderWithProviders } from '../../../../../test-utils'
import OrderForm from './index'

test('the form values of the onSubmitHandler', async () => {
  const onSubmitHandlerSpy = jest.fn()

  const { getByLabelText, getByText } = renderWithProviders(
    <OrderForm
      services={[
        {
          organization: 'organization-a',
          service: 'service-a',
        },
      ]}
      onSubmitHandler={onSubmitHandlerSpy}
    />,
  )

  userEvent.type(getByLabelText(/Order description/), 'my-description')
  userEvent.type(getByLabelText(/Reference/), 'my-reference')
  userEvent.type(getByLabelText(/Public key PEM/), 'my-public-key-pem')
  userEvent.type(getByLabelText(/Delegated organization/), 'my-delegatee')
  fireEvent.change(getByLabelText(/Valid from/), {
    target: { value: '2021-01-01' },
  })
  fireEvent.change(getByLabelText(/Valid until/), {
    target: { value: '2021-01-31' },
  })
  await selectEvent.select(getByLabelText(/Services/), /service-a/)

  userEvent.click(getByText('Add order'))

  await waitFor(() =>
    expect(onSubmitHandlerSpy).toHaveBeenCalledWith({
      delegatee: 'my-delegatee',
      description: 'my-description',
      publicKeyPEM: 'my-public-key-pem',
      reference: 'my-reference',
      services: [
        {
          organization: 'organization-a',
          service: 'service-a',
        },
      ],
      validFrom: new Date('2021-01-01T00:00:00.000Z'),
      validUntil: new Date('2021-01-31T00:00:00.000Z'),
    }),
  )
})
