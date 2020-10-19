// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../models/OutgoingAccessRequestModel'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'

test('creating an outgoing access request', async () => {
  const accessRequestRepository = {
    createAccessRequest: jest.fn().mockResolvedValue({
      id: '42',
      organizationName: 'organization-name',
      serviceName: 'service-name',
      state: ACCESS_REQUEST_STATES.CREATED,
      createdAt: '2020-10-07T13:01:11.288349Z',
      updatedAt: null,
    }),
  }

  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
    accessRequestRepository: accessRequestRepository,
  })

  const outgoingAccessRequest = await outgoingAccessRequestStore.create({
    organizationName: 'organization-name',
    serviceName: 'service-name',
  })

  expect(accessRequestRepository.createAccessRequest).toHaveBeenCalledWith({
    organizationName: 'organization-name',
    serviceName: 'service-name',
  })

  expect(outgoingAccessRequest).toEqual(
    new OutgoingAccessRequestModel({
      accessRequestData: {
        id: '42',
        organizationName: 'organization-name',
        serviceName: 'service-name',
        state: ACCESS_REQUEST_STATES.CREATED,
        createdAt: '2020-10-07T13:01:11.288349Z',
        updatedAt: null,
      },
    }),
  )
})
