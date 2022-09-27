// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import OutwayModel from './OutwayModel'

test('creating an Outway instance', () => {
  const model = new OutwayModel({
    outwayData: {
      name: 'my-outway',
      ipAddress: '127.0.0.1',
      publicKeyPem: 'test-pem',
      version: 'v0.0.42',
    },
  })

  expect(model.name).toEqual('my-outway')
  expect(model.ipAddress).toBe('127.0.0.1')
  expect(model.publicKeyPem).toBe('test-pem')
  expect(model.version).toBe('v0.0.42')
})
