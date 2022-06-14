// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import TransactionLogModel, { DIRECTION_IN } from './TransactionLogModel'

test('initialize and update the transactionLog', async () => {
  const transactionLogModel = new TransactionLogModel({
    transactionLogsStore: {},
    transactionLogData: {
      direction: DIRECTION_IN,
      source: {
        serialNumber: '00000000000000000001',
        name: 'Organization One',
      },
      destination: {
        serialNumber: '00000000000000000002',
        name: 'Organization Two',
      },
      serviceName: 'test-service',
      order: {
        delegator: {
          serialNumber: '00000000000000000003',
          name: 'Organization Three',
        },
        reference: 'test-reference',
      },
      createdAt: new Date(),
    },
  })

  transactionLogModel.update({
    service: {
      name: 'test-service-2',
    },
  })

  expect(transactionLogModel.serviceName).toBe('test-service-2')
  expect(transactionLogModel.source.name).toBe('Organization One')
  expect(transactionLogModel.destination.name).toBe('Organization Two')
  expect(transactionLogModel.order.delegator.name).toBe('Organization Three')
})
