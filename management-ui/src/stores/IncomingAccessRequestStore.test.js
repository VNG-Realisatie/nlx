// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import IncomingAccessRequestStore from './OutgoingAccessRequestStore'

test('initializing', async () => {
  const incomingAccessRequestStore = new IncomingAccessRequestStore({})
  expect(incomingAccessRequestStore).toBeTruthy()
})
