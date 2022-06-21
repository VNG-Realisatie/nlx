// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import TransactionLogModel, {
  DIRECTION_IN,
} from '../../../stores/models/TransactionLogModel'
import TransactionLogRow from './index'

test('transaction row should render expected data', () => {
  const model = new TransactionLogModel({
    transactionLogData: {
      source: {
        serialNumber: '00000000000000000001',
        name: 'Organization One',
      },
      destination: {
        serialNumber: '00000000000000000002',
        name: 'Organization Two',
      },
      service: {
        name: 'my-service',
      },
      direction: DIRECTION_IN,
      createdAt: new Date(),
    },
  })

  const { queryByText, rerender } = renderWithProviders(
    <table>
      <tbody>
        <MemoryRouter>
          <TransactionLogRow transactionLog={model} />
        </MemoryRouter>
      </tbody>
    </table>,
  )

  expect(queryByText('my-service')).toBeInTheDocument()
  expect(queryByText('Organization One')).toBeInTheDocument()

  model.update({
    order: {
      reference: 'ref-1',
      delegator: {
        serialNumber: '00000000000000000002',
        name: 'Organization Two',
      },
    },
  })

  rerender(
    <table>
      <tbody>
        <MemoryRouter>
          <TransactionLogRow transactionLog={model} />
        </MemoryRouter>
      </tbody>
    </table>,
  )

  expect(
    queryByText(`Organization One On behalf of Organization Two`),
  ).toBeInTheDocument()
})
