// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AuditLogModel, { ACTION_LOGIN_SUCCESS } from './AuditLogModel'

test('creating AuditLog instance', () => {
  const model = new AuditLogModel({
    auditLogData: {
      id: '42',
      action: ACTION_LOGIN_SUCCESS,
      user: 'John Doe',
      createdAt: '2020-10-01T12:00:00Z',
      services: [
        {
          organization: 'Gemeente Amsterdam',
          service: 'vakantieverhuur',
        },
      ],
      data: {
        delegatee: {
          serialNumber: '00000000000000000001',
          name: 'Kadaster',
        },
        reference: '030394AB',
        inwayName: 'my-inway',
      },
    },
  })

  expect(model.id).toEqual('42')
  expect(model.action).toBe(ACTION_LOGIN_SUCCESS)
  expect(model.user).toBe('John Doe')
  expect(model.createdAt).toEqual(new Date('2020-10-01T12:00:00Z'))
  expect(model.services).toStrictEqual([
    {
      organization: 'Gemeente Amsterdam',
      service: 'vakantieverhuur',
    },
  ])
  expect(model.data.delegatee.serialNumber).toEqual('00000000000000000001')
  expect(model.data.delegatee.name).toEqual('Kadaster')
  expect(model.data.reference).toEqual('030394AB')
  expect(model.data.inwayName).toEqual('my-inway')
})

test('organization name is empty', () => {
  const auditLog = new AuditLogModel({
    auditLogData: {
      data: {
        delegatee: {
          serialNumber: '00000000000000000001',
          name: '',
        },
      },
    },
  })

  expect(auditLog.data.delegatee.serialNumber).toEqual('00000000000000000001')
  expect(auditLog.data.delegatee.name).toEqual('00000000000000000001')
})
