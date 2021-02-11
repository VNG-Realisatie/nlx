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

  managementApiClient.managementListAuditLogs = jest.fn().mockResolvedValue({
    auditLogs: [
      {
        id: '42',
        user: 'John Doe',
      },
    ],
  })

  const auditLogStore = new AuditLogStore({
    managementApiClient,
  })

  await auditLogStore.fetchAll()
  expect(auditLogStore.auditLogs).toHaveLength(1)
  expect(auditLogStore.auditLogs[0]).toBeInstanceOf(AuditLogModel)

  const initialAuditLog = auditLogStore.auditLogs[0]
  const updatedAuditLog = await auditLogStore.updateFromServer({
    id: '42',
    user: 'Jane Doe',
  })

  expect(initialAuditLog).toBe(updatedAuditLog)
  expect(auditLogStore.auditLogs).toHaveLength(1)
  expect(updatedAuditLog.user).toEqual('Jane Doe')
})
