// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import TransactionLogModel, { DIRECTION_IN } from './TransactionLogModel'

let transactionLogData

beforeEach(() => {
  transactionLogData = {
    direction: DIRECTION_IN,
    source: {
      serialNumber: '0001',
    },
    destination: {
      serialNumber: '0002',
    },
    serviceName: 'test-service',
    order: {
      delegator: '0003',
      reference: 'test-reference',
    },
    createdAt: new Date(),
  }
})

afterEach(() => {
  jest.restoreAllMocks()
})

test('initialize and update the transactionLog', async () => {
  const transactionLogModel = new TransactionLogModel({
    transactionLogsStore: {},
    transactionLogData,
  })

  transactionLogData.service = { name: 'test-service-2' }

  transactionLogModel.update(transactionLogData)

  expect(transactionLogModel.serviceName).toBe('test-service-2')
})
