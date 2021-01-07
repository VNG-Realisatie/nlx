// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import ServiceModel from '../stores/models/ServiceModel'
import AccessGrantModel from '../stores/models/AccessGrantModel'
import { ManagementApi } from '../api'
import AccessGrantStore from './AccessGrantStore'

test('initializing the store', () => {
  const accessGrantStore = new AccessGrantStore({
    managementApiClient: new ManagementApi(),
  })

  expect(accessGrantStore.accessGrants.size).toEqual(0)
})

test('fetching, getting and updating from server', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListAccessGrantsForService = jest
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

  const service = new ServiceModel({
    serviceData: { name: 'Service' },
  })
  await accessGrantStore.fetchForService(service)

  expect(
    managementApiClient.managementListAccessGrantsForService,
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
  const managementApiClient = new ManagementApi()
  const accessGrantStore = new AccessGrantStore({
    managementApiClient,
  })

  managementApiClient.managementRevokeAccessGrant = jest
    .fn()
    .mockResolvedValue()

  jest.spyOn(accessGrantStore, 'fetchForService').mockResolvedValue()

  await accessGrantStore.revokeAccessGrant({
    organizationName: 'Organization',
    serviceName: 'Service',
    id: 's1',
  })

  expect(managementApiClient.managementRevokeAccessGrant).toHaveBeenCalledWith({
    organizationName: 'Organization',
    serviceName: 'Service',
    accessGrantID: 's1',
  })
  expect(accessGrantStore.fetchForService).toHaveBeenCalledWith({
    name: 'Service',
  })
})
