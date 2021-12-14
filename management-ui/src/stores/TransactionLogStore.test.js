// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { TXLogApi } from '../api'
import TransactionLogModel from './models/TransactionLogModel'
import TransactionLogStore from './TransactionLogStore'

test('initializing the store', () => {
  const txLogApiClient = new TXLogApi()
  const transactionLogStore = new TransactionLogStore({
    txLogApiClient: txLogApiClient,
  })

  expect(transactionLogStore.transactionLogs).toEqual([])
})

test('fetching, getting and updating from server', async () => {
  configure({ safeDescriptors: false })

  const txLogApiClient = new TXLogApi()
  txLogApiClient.tXLogListRecords = jest
    .fn()
    .mockResolvedValueOnce({
      records: [
        {
          direction: 'IN',
          source: {
            serialNumber: '00000000000000000001',
          },
          destination: {
            serialNumber: '00000000000000000002',
          },
          service: {
            name: 'my-service',
          },
        },
      ],
    })
    .mockResolvedValue({
      records: [
        {
          direction: 'IN',
          source: {
            serialNumber: '00000000000000000001',
          },
          destination: {
            serialNumber: '00000000000000000002',
          },
          service: {
            name: 'my-service3',
          },
        },
        {
          direction: 'OUT',
          source: {
            serialNumber: '00000000000000000001',
          },
          destination: {
            serialNumber: '00000000000000000002',
          },
          service: {
            name: 'my-service2',
          },
        },
      ],
    })

  const transactionLogStore = new TransactionLogStore({
    txLogApiClient,
  })

  await transactionLogStore.fetchAll()
  expect(transactionLogStore.transactionLogs).toHaveLength(1)
  const intialTransactionLog = transactionLogStore.transactionLogs[0]
  expect(intialTransactionLog).toBeInstanceOf(TransactionLogModel)
  expect(intialTransactionLog.serviceName).toEqual('my-service')

  await transactionLogStore.fetchAll()

  expect(transactionLogStore.transactionLogs).toHaveLength(2)
  const secondTransactionLog = transactionLogStore.transactionLogs[1]
  expect(secondTransactionLog.serviceName).toEqual('my-service2')
})
