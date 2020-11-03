// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import ServiceModel from '../models/ServiceModel'
import IncomingAccessRequestModel from '../models/IncomingAccessRequestModel'
import IncomingAccessRequestsStore from './IncomingAccessRequestsStore'

test('fetching, getting and updating from server (integration test)', async () => {
  const service = new ServiceModel({ serviceData: { name: 'Service' } })

  const incomingAccessRequestData = {
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    state: 'CREATED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:01Z',
  }
  const fetchByServiceName = jest
    .fn()
    .mockResolvedValue([incomingAccessRequestData])
  const incomingAccessRequestStore = new IncomingAccessRequestsStore({
    accessRequestRepository: {
      fetchByServiceName,
    },
  })

  expect(incomingAccessRequestStore.incomingAccessRequests.size).toEqual(0)
  expect(incomingAccessRequestStore.updateFromServer()).toBeNull()

  await incomingAccessRequestStore.fetchForService(service)

  expect(fetchByServiceName).toHaveBeenCalledWith(service.name)
  expect(incomingAccessRequestStore.incomingAccessRequests.size).toEqual(1)

  const accessRequestsForService = incomingAccessRequestStore.getForService(
    service,
  )
  expect(accessRequestsForService).toHaveLength(1)
  expect(accessRequestsForService[0]).toBeInstanceOf(IncomingAccessRequestModel)

  const updatedAccessRequest = await incomingAccessRequestStore.updateFromServer(
    {
      id: 'abcd',
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'RECEIVED',
      createdAt: '2020-10-01T12:00:00Z',
      updatedAt: '2020-10-01T12:00:10Z',
    },
  )

  expect(incomingAccessRequestStore.incomingAccessRequests.size).toEqual(1)
  expect(updatedAccessRequest).toBeInstanceOf(IncomingAccessRequestModel)
  expect(updatedAccessRequest.state).toBe('RECEIVED')
})
