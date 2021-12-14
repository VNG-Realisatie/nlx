// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import { DIRECTION_IN } from '../../../stores/models/TransactionLogModel'
import TransactionLogRow from './index'

test('transaction row should render expected data', () => {
  const transactionLogRecord = {
    source: {
      serialNumber: '00000000000000000001',
    },
    destination: {
      serialNumber: '00000000000000000002',
    },
    serviceName: 'my-service',
    direction: DIRECTION_IN,
    createdAt: new Date(),
  }

  const { queryByText, rerender } = renderWithProviders(
    <table>
      <tbody>
        <TransactionLogRow transactionLog={transactionLogRecord} />
      </tbody>
    </table>,
  )

  expect(queryByText('my-service')).toBeInTheDocument()
  expect(queryByText('00000000000000000001')).toBeInTheDocument()

  const transactionLogWithOrder = Object.assign({}, transactionLogRecord, {
    order: {
      reference: 'ref-1',
      delegator: '00000000000000000002',
    },
  })

  rerender(
    <table>
      <tbody>
        <TransactionLogRow transactionLog={transactionLogWithOrder} />
      </tbody>
    </table>,
  )

  expect(
    queryByText(`00000000000000000001 On behalf of 00000000000000000002`),
  ).toBeInTheDocument()
})
