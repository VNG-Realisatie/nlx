// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { TXLogServiceApi, ManagementServiceApi } from '../api'
import TransactionLogModel from './models/TransactionLogModel'
import TransactionLogStore from './TransactionLogStore'

test('initializing the store', () => {
  const txLogApiClient = new TXLogServiceApi()
  const managementApiClient = new ManagementServiceApi()
  const transactionLogStore = new TransactionLogStore({
    txLogApiClient: txLogApiClient,
    managementApiClient: managementApiClient,
  })

  expect(transactionLogStore.transactionLogs).toEqual([])
})

test('fetching, getting and updating from server', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceIsTXLogEnabled = jest
    .fn()
    .mockResolvedValue({
      enabled: true,
    })

  const txLogApiClient = new TXLogServiceApi()
  txLogApiClient.tXLogServiceListRecords = jest
    .fn()
    .mockResolvedValueOnce({
      records: [
        {
          transactionId: '2d37d10f3b6515b4075278877629d116',
          direction: 'DIRECTION_IN',
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
        },
      ],
    })
    .mockResolvedValue({
      records: [
        {
          transactionId: '2d37d10f3b6515b4075278877629d116',
          direction: 'DIRECTION_IN',
          source: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
          destination: {
            serialNumber: '00000000000000000002',
            name: 'Organization Two',
          },
          service: {
            name: 'my-service3',
          },
        },
        {
          transactionId: '70c5cb7ef23af6a416b2b47a377dd39f',
          direction: 'DIRECTION_OUT',
          source: {
            serialNumber: '00000000000000000001',
            name: 'Organization One',
          },
          destination: {
            serialNumber: '00000000000000000002',
            name: 'Organization Two',
          },
          service: {
            name: 'my-service2',
          },
        },
      ],
    })

  const transactionLogStore = new TransactionLogStore({
    txLogApiClient,
    managementApiClient,
  })

  await transactionLogStore.fetchAll()
  expect(transactionLogStore.transactionLogs).toHaveLength(1)
  const intialTransactionLog = transactionLogStore.transactionLogs[0]
  expect(intialTransactionLog).toBeInstanceOf(TransactionLogModel)
  expect(intialTransactionLog.serviceName).toEqual('my-service')
  expect(intialTransactionLog.transactionId).toEqual(
    '2d37d10f3b6515b4075278877629d116',
  )

  await transactionLogStore.fetchAll()

  expect(transactionLogStore.transactionLogs).toHaveLength(2)
  const secondTransactionLog = transactionLogStore.transactionLogs[1]
  expect(secondTransactionLog.serviceName).toEqual('my-service2')
  expect(secondTransactionLog.transactionId).toEqual(
    '70c5cb7ef23af6a416b2b47a377dd39f',
  )
})
