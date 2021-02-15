// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import AuditLogStore from './AuditLogStore'
import AuditLogModel from './models/AuditLogModel'

test('initializing the store', () => {
  const auditLogStore = new AuditLogStore({
    managementApiClient: new ManagementApi(),
  })

  expect(auditLogStore.auditLogs).toEqual([])
})

test('fetching, getting and updating from server', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListAuditLogs = jest
    .fn()
    .mockResolvedValueOnce({
      auditLogs: [
        {
          id: '42',
          user: 'John Doe',
        },
      ],
    })
    .mockResolvedValue({
      auditLogs: [
        {
          id: '41',
          user: 'Jane Doe',
        },
        {
          id: '42',
          user: 'Peter Doe',
        },
      ],
    })

  const auditLogStore = new AuditLogStore({
    managementApiClient,
  })

  await auditLogStore.fetchAll()
  expect(auditLogStore.auditLogs).toHaveLength(1)
  const initialAuditLog = auditLogStore.auditLogs[0]
  expect(initialAuditLog).toBeInstanceOf(AuditLogModel)

  await auditLogStore.fetchAll()

  expect(auditLogStore.auditLogs).toHaveLength(2)
  const updatedAuditLog = auditLogStore.auditLogs[1]
  expect(initialAuditLog).toBe(updatedAuditLog)
  expect(updatedAuditLog.user).toEqual('Peter Doe')
})
