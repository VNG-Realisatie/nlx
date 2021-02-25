// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import IncomingAccessRequestModel, {
  STATES,
} from '../stores/models/IncomingAccessRequestModel'
import { ManagementApi } from '../api'
import IncomingAccessRequestsStore from './IncomingAccessRequestsStore'
import { RootStore } from './index'

test('initializing the store', () => {
  const incomingAccessRequestStore = new IncomingAccessRequestsStore({})
  expect(incomingAccessRequestStore.incomingAccessRequests.size).toEqual(0)
})

test('fetching, getting and updating from server', async () => {
  const service = { name: 'Service' }

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
  configure({ safeDescriptors: false })
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

test('fetching for a service should update existing in-memory models instead of recreating them', async () => {
  const managementApiClient = new ManagementApi()

  const incomingAccessRequestStore = new IncomingAccessRequestsStore({
    managementApiClient,
  })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValueOnce({
      accessRequests: [
        {
          id: 'ar-1',
          serviceName: 'service-a',
          organizationName: 'organization-a',
          state: STATES.CREATED,
        },
        {
          id: 'ar-2',
          serviceName: 'service-a',
          organizationName: 'organization-b',
          state: STATES.CREATED,
        },
      ],
    })
    .mockResolvedValue({
      accessRequests: [
        {
          id: 'ar-1',
          serviceName: 'service-a',
          organizationName: 'organization-a',
          state: STATES.CREATED,
        },
        {
          id: 'ar-2',
          serviceName: 'service-a',
          organizationName: 'organization-b',
          state: STATES.CREATED,
        },
      ],
    })

  await incomingAccessRequestStore.fetchForService({ name: 'service-a' })
  const initialAccessRequests = incomingAccessRequestStore.getForService({
    name: 'service-a',
  })

  await incomingAccessRequestStore.fetchForService({ name: 'service-a' })
  const newAccessRequests = incomingAccessRequestStore.getForService({
    name: 'service-a',
  })

  expect(initialAccessRequests[0]).toBe(newAccessRequests[0])
})

describe('have the access requests been changed for a service', () => {
  let managementApiClient
  let rootStore
  let servicesStore
  let incomingAccessRequestStore

  const ACCESS_REQUEST_ONE = {
    id: 'ar-1',
    serviceName: 'service-a',
    organizationName: 'organization-a',
    state: STATES.CREATED,
  }

  const ACCESS_REQUEST_TWO = {
    id: 'ar-2',
    serviceName: 'service-a',
    organizationName: 'organization-a',
    state: STATES.CREATED,
  }

  beforeEach(async () => {
    managementApiClient = new ManagementApi()

    managementApiClient.managementGetService = jest.fn().mockResolvedValue({
      name: 'service-a',
    })

    rootStore = new RootStore({
      managementApiClient,
    })

    rootStore.accessGrantStore.fetchForService = jest.fn().mockResolvedValue()

    incomingAccessRequestStore = rootStore.incomingAccessRequestsStore
    servicesStore = rootStore.servicesStore
  })

  it('should indicate changed if the number of access requests have changed', async () => {
    managementApiClient.managementListIncomingAccessRequest = jest
      .fn()
      .mockResolvedValueOnce({ accessRequests: [] })
      .mockResolvedValue({
        accessRequests: [ACCESS_REQUEST_ONE],
      })

    await servicesStore.fetch({ name: 'service-a' })
    const service = servicesStore.getService('service-a')

    expect(
      await incomingAccessRequestStore.haveChangedForService(service),
    ).toEqual(true)
  })

  it('should indicate changed if the changed access request has a different id', async () => {
    managementApiClient.managementListIncomingAccessRequest = jest
      .fn()
      .mockResolvedValueOnce({
        accessRequests: [ACCESS_REQUEST_ONE],
      })
      .mockResolvedValue({
        accessRequests: [ACCESS_REQUEST_TWO],
      })

    await servicesStore.fetch({ name: 'service-a' })
    const service = servicesStore.getService('service-a')

    expect(
      await incomingAccessRequestStore.haveChangedForService(service),
    ).toEqual(true)
  })

  it('should not indicate changed if the latest access request is the same', async () => {
    managementApiClient.managementListIncomingAccessRequest = jest
      .fn()
      .mockResolvedValue({
        accessRequests: [ACCESS_REQUEST_ONE],
      })

    await servicesStore.fetch({ name: 'service-a' })
    const service = servicesStore.getService('service-a')

    expect(
      await incomingAccessRequestStore.haveChangedForService(service),
    ).toEqual(false)
  })
})
