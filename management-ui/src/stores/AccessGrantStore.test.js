// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import AccessGrantModel from '../stores/models/AccessGrantModel'
import { ManagementServiceApi } from '../api'
import AccessGrantStore from './AccessGrantStore'
import { RootStore } from './index'

test('initializing the store', () => {
  const accessGrantStore = new AccessGrantStore({
    managementApiClient: new ManagementServiceApi(),
  })

  expect(accessGrantStore.accessGrants.size).toEqual(0)
})

test('fetching, getting and updating from server', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [
        {
          id: 'abcd',
          organizationName: 'Organization',
          serviceName: 'Service',
          createdAt: '2020-10-01',
          revokedAt: null,
        },
      ],
    })

  const accessGrantStore = new AccessGrantStore({
    managementApiClient,
  })

  const service = { name: 'Service' }
  await accessGrantStore.fetchForService(service)

  expect(
    managementApiClient.managementServiceListAccessGrantsForService,
  ).toHaveBeenCalledWith({ serviceName: 'Service' })

  expect(accessGrantStore.accessGrants.size).toEqual(1)

  const accessGrantsForService = accessGrantStore.getForService(service)
  expect(accessGrantsForService).toHaveLength(1)
  expect(accessGrantsForService[0]).toBeInstanceOf(AccessGrantModel)

  const updatedAccessGrant = await accessGrantStore.updateFromServer({
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: '2020-10-02',
  })

  expect(accessGrantStore.accessGrants.size).toEqual(1)
  expect(updatedAccessGrant).toBeInstanceOf(AccessGrantModel)
  expect(updatedAccessGrant.revokedAt).toEqual(new Date('2020-10-02'))
})

test('revoking an access grant', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementServiceApi()
  const accessGrantStore = new AccessGrantStore({
    managementApiClient,
  })

  managementApiClient.managementServiceRevokeAccessGrant = jest
    .fn()
    .mockResolvedValue()

  jest.spyOn(accessGrantStore, 'fetchForService').mockResolvedValue()

  await accessGrantStore.revokeAccessGrant({
    serviceName: 'Service',
    id: 's1',
  })

  expect(
    managementApiClient.managementServiceRevokeAccessGrant,
  ).toHaveBeenCalledWith({
    accessGrantId: 's1',
  })
  expect(accessGrantStore.fetchForService).toHaveBeenCalledWith({
    name: 'Service',
  })
})

test('fetching for a service should update existing in-memory models instead of recreating them', async () => {
  const managementApiClient = new ManagementServiceApi()

  const accessGrantStore = new AccessGrantStore({
    managementApiClient,
  })

  accessGrantStore.updateFromServer([
    {
      id: 'ag-1',
      serviceName: 'service-a',
      organizationName: 'organization-a',
    },
    {
      id: 'ag-2',
      serviceName: 'service-a',
      organizationName: 'organization-b',
    },
  ])

  managementApiClient.managementServiceListAccessGrantsForService = jest
    .fn()
    .mockResolvedValue({
      accessGrants: [
        {
          id: 'ag-1',
          serviceName: 'service-a',
          organizationName: 'organization-a',
        },
        {
          id: 'ag-2',
          serviceName: 'service-a',
          organizationName: 'organization-b',
        },
      ],
    })

  await accessGrantStore.fetchForService({ name: 'service-a' })
  const initialAccessGrants = accessGrantStore.getForService({
    name: 'service-a',
  })

  await accessGrantStore.fetchForService({ name: 'service-a' })
  const newAccessGrants = accessGrantStore.getForService({
    name: 'service-a',
  })

  expect(initialAccessGrants[0]).toBe(newAccessGrants[0])
})

describe('have the access grants been changed for a service', () => {
  let managementApiClient
  let rootStore
  let servicesStore
  let accessGrantStore

  const ACCESS_GRANT_ONE = {
    id: 'ar-1',
    serviceName: 'service-a',
    organizationName: 'organization-a',
    revokedAt: null,
  }

  const ACCESS_GRANT_TWO = {
    id: 'ar-2',
    serviceName: 'service-a',
    organizationName: 'organization-a',
    revokedAt: undefined,
  }

  beforeEach(async () => {
    managementApiClient = new ManagementServiceApi()

    managementApiClient.managementServiceGetService = jest
      .fn()
      .mockResolvedValue({
        name: 'service-a',
      })

    rootStore = new RootStore({
      managementApiClient,
    })

    rootStore.incomingAccessRequestsStore.fetchForService = jest
      .fn()
      .mockResolvedValue()

    accessGrantStore = rootStore.accessGrantStore
    servicesStore = rootStore.servicesStore
  })

  it('should indicate changed if the number of access grants have changed', async () => {
    managementApiClient.managementServiceListAccessGrantsForService = jest
      .fn()
      .mockResolvedValueOnce({ accessGrants: [] })
      .mockResolvedValue({
        accessGrants: [ACCESS_GRANT_ONE],
      })

    await servicesStore.fetch({ name: 'service-a' })
    const service = servicesStore.getByName('service-a')

    expect(await accessGrantStore.haveChangedForService(service)).toEqual(true)
  })

  it('should indicate changed if the changed access request has a different id', async () => {
    managementApiClient.managementServiceListAccessGrantsForService = jest
      .fn()
      .mockResolvedValueOnce({
        accessGrants: [ACCESS_GRANT_ONE],
      })
      .mockResolvedValue({
        accessGrants: [ACCESS_GRANT_TWO],
      })

    await servicesStore.fetch({ name: 'service-a' })
    const service = servicesStore.getByName('service-a')

    expect(await accessGrantStore.haveChangedForService(service)).toEqual(true)
  })

  it('should not indicate changed if the latest access request is the same', async () => {
    managementApiClient.managementServiceListAccessGrantsForService = jest
      .fn()
      .mockResolvedValue({
        accessGrants: [ACCESS_GRANT_ONE],
      })

    await servicesStore.fetch({ name: 'service-a' })
    const service = servicesStore.getByName('service-a')

    expect(await accessGrantStore.haveChangedForService(service)).toEqual(false)
  })
})
