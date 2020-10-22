// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import IncomingAccessRequestsStore from './IncomingAccessRequestsStore'

test('initializing', async () => {
  const incomingAccessRequestsStore = new IncomingAccessRequestsStore({})
  expect(incomingAccessRequestsStore).toBeTruthy()
})
