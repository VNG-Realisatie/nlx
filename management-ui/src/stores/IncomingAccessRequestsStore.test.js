// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import ServiceModel from '../stores/models/ServiceModel'
import IncomingAccessRequestModel from '../stores/models/IncomingAccessRequestModel'
import { ManagementApi } from '../api'
import IncomingAccessRequestsStore from './IncomingAccessRequestsStore'

test('initializing the store', () => {
  const incomingAccessRequestStore = new IncomingAccessRequestsStore({})
  expect(incomingAccessRequestStore.incomingAccessRequests.size).toEqual(0)
})

test('fetching, getting and updating from server', async () => {
  const service = new ServiceModel({ serviceData: { name: 'Service' } })

  const managementApiClient = new ManagementApi()
  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue({
      accessRequests: [
        {
          id: 'abcd',
          organizationName: 'Organization',
          serviceName: 'Service',
          state: 'RECEIVED',
          createdAt: '2020-10-01T12:00:00Z',
          updatedAt: '2020-10-01T12:00:10Z',
        },
      ],
    })
    .mockResolvedValueOnce({
      accessRequests: [
        {
          id: 'abcd',
          organizationName: 'Organization',
          serviceName: 'Service',
          state: 'CREATED',
          createdAt: '2020-10-01T12:00:00Z',
          updatedAt: '2020-10-01T12:00:01Z',
        },
      ],
    })

  const incomingAccessRequestStore = new IncomingAccessRequestsStore({
    managementApiClient,
  })

  expect(incomingAccessRequestStore.updateFromServer()).toBeNull()

  await incomingAccessRequestStore.fetchForService(service)

  expect(
    managementApiClient.managementListIncomingAccessRequest,
  ).toHaveBeenCalledWith({ serviceName: 'Service' })
  expect(incomingAccessRequestStore.incomingAccessRequests.size).toEqual(1)

  const accessRequestsForService = incomingAccessRequestStore.getForService(
    service,
  )
  expect(accessRequestsForService).toHaveLength(1)
  expect(accessRequestsForService[0]).toBeInstanceOf(IncomingAccessRequestModel)

  await incomingAccessRequestStore.fetchForService(service)
  const updatedAccessRequests = incomingAccessRequestStore.getForService(
    service,
  )

  expect(incomingAccessRequestStore.incomingAccessRequests.size).toEqual(1)
  expect(updatedAccessRequests[0]).toBeInstanceOf(IncomingAccessRequestModel)
  expect(updatedAccessRequests[0].state).toBe('RECEIVED')
})

test('approving an access request', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementApproveIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue()

  const incomingAccessRequestStore = new IncomingAccessRequestsStore({
    managementApiClient,
  })

  const fetchForServiceSpy = jest
    .spyOn(incomingAccessRequestStore, 'fetchForService')
    .mockResolvedValue()

  await incomingAccessRequestStore.approveAccessRequest({
    serviceName: 'Service',
    id: 's1',
  })

  expect(
    managementApiClient.managementApproveIncomingAccessRequest,
  ).toHaveBeenCalled()
  expect(fetchForServiceSpy).toHaveBeenCalledWith({ name: 'Service' })
})

test('rejecting an access request', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementRejectIncomingAccessRequest = jest
    .fn()
    .mockResolvedValue()

  const incomingAccessRequestStore = new IncomingAccessRequestsStore({
    managementApiClient,
  })

  const fetchForServiceSpy = jest
    .spyOn(incomingAccessRequestStore, 'fetchForService')
    .mockResolvedValue()

  await incomingAccessRequestStore.rejectAccessRequest({
    serviceName: 'Service',
    id: 's1',
  })

  expect(
    managementApiClient.managementRejectIncomingAccessRequest,
  ).toHaveBeenCalled()
  expect(fetchForServiceSpy).toHaveBeenCalledWith({ name: 'Service' })
})
