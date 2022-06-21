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

test('organization name is empty', () => {
  const model = new TransactionLogModel({
    transactionLogStore: {},
    transactionLogData: {
      source: {
        serialNumber: '00000000000000000001',
        name: '',
      },
      destination: {
        serialNumber: '00000000000000000002',
        name: '',
      },
      order: {
        delegator: {
          serialNumber: '00000000000000000003',
          name: '',
        },
      },
    },
  })

  expect(model.source.name).toBe('00000000000000000001')
  expect(model.source.serialNumber).toBe('00000000000000000001')

  expect(model.destination.name).toBe('00000000000000000002')
  expect(model.destination.serialNumber).toBe('00000000000000000002')

  expect(model.order.delegator.name).toBe('00000000000000000003')
  expect(model.order.delegator.serialNumber).toBe('00000000000000000003')
})
