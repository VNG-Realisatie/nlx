// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import TransactionLogModel, {
  DIRECTION_IN,
  DIRECTION_OUT,
} from '../../../stores/models/TransactionLogModel'
import TransactionLogRow from './index'

test('transaction row direction IN should render expected data', () => {
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
      order: {
        delegator: {
          serialNumber: '',
          name: '',
        },
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

test('transaction row with direction OUT should render expected data', () => {
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
      order: {
        delegator: {
          serialNumber: '',
          name: '',
        },
      },
      direction: DIRECTION_OUT,
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
  expect(queryByText('Organization Two')).toBeInTheDocument()

  model.update({
    order: {
      reference: 'ref-1',
      delegator: {
        serialNumber: '00000000000000000003',
        name: 'Organization Three',
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
    queryByText(`Organization Two On behalf of Organization Three`),
  ).toBeInTheDocument()
})
