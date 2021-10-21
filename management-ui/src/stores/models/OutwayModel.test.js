// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import OutwayModel from './OutwayModel'

test('creating an Outway instance', () => {
  const model = new OutwayModel({
    outwayData: {
      name: 'my-outway',
      version: 'v0.0.42',
    },
  })

  expect(model.name).toEqual('my-outway')
  expect(model.version).toBe('v0.0.42')
})
