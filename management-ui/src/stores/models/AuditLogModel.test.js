// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AuditLogModel, { ACTION_LOGIN_SUCCESS } from './AuditLogModel'

test('creating AuditLog instance', () => {
  const auditLog = new AuditLogModel({
    auditLogData: {
      id: '42',
      action: ACTION_LOGIN_SUCCESS,
      user: 'John Doe',
      createdAt: '2020-10-01T12:00:00Z',
      organization: 'My BigCorp',
    },
  })

  expect(auditLog.id).toEqual('42')
  expect(auditLog.action).toBe(ACTION_LOGIN_SUCCESS)
  expect(auditLog.user).toBe('John Doe')
  expect(auditLog.createdAt).toEqual(new Date('2020-10-01T12:00:00Z'))
  expect(auditLog.organization).toBe('My BigCorp')
})
