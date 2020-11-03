// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import ServiceModel from '../models/ServiceModel'
import AccessGrantModel from '../models/AccessGrantModel'
import AccessGrantStore from './AccessGrantStore'

test('fetching, getting and updating from server (integration test)', async () => {
  const service = new ServiceModel({ serviceData: { name: 'Service' } })

  const accessGrantData = {
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    createdAt: '2020-10-01',
    revokedAt: null,
  }
  const fetchByServiceName = jest.fn().mockResolvedValue([accessGrantData])
  const accessGrantStore = new AccessGrantStore({
    accessGrantRepository: {
      fetchByServiceName,
    },
  })

  expect(accessGrantStore.accessGrants.size).toEqual(0)
  expect(accessGrantStore.updateFromServer()).toBeNull()

  await accessGrantStore.fetchForService(service)

  expect(fetchByServiceName).toHaveBeenCalledWith(service.name)
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
