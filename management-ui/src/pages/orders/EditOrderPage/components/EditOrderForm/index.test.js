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
  const onSubmitHandlerMock = jest.fn()

  const testServices = [
    {
      organization: {
        serialNumber: '00000000000000000001',
        name: 'organization-a',
      },
      service: 'service-a',
    },
  ]

  const { getByLabelText, getByText } = renderWithProviders(
    <OrderForm
      order={{
        description: 'old',
        publicKeyPEM: 'old',
        validFrom: '2020-12-01',
        validUntil: '2021-12-01',
        services: testServices,
      }}
      services={testServices}
      onSubmitHandler={onSubmitHandlerMock}
    />,
  )

  fireEvent.change(getByLabelText(/Order description/), {
    target: { value: '' },
  })
  userEvent.type(getByLabelText(/Order description/), 'my-description')
  fireEvent.change(getByLabelText(/Public key PEM/), { target: { value: '' } })
  userEvent.type(getByLabelText(/Public key PEM/), 'my-public-key-pem')

  fireEvent.change(getByLabelText(/Valid from/), {
    target: { value: '2021-01-01' },
  })
  fireEvent.change(getByLabelText(/Valid until/), {
    target: { value: '2021-01-31' },
  })
  await selectEvent.select(getByLabelText(/Services/), /service-a/)

  userEvent.click(getByText(/Update order/))

  await waitFor(() =>
    expect(onSubmitHandlerMock).toHaveBeenCalledWith({
      description: 'my-description',
      publicKeyPEM: 'my-public-key-pem',
      services: [
        {
          organization: {
            serialNumber: '00000000000000000001',
            name: 'organization-a',
          },
          service: 'service-a',
        },
      ],
      validFrom: new Date('2021-01-01T00:00:00.000Z'),
      validUntil: new Date('2021-01-31T00:00:00.000Z'),
    }),
  )
})
