// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../models/OutgoingAccessRequestModel'
import { ManagementApi } from '../api'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'

test('creating an outgoing access request', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementCreateAccessRequest = jest
    .fn()
    .mockResolvedValue({
      id: '42',
      organizationName: 'organization-name',
      serviceName: 'service-name',
      state: ACCESS_REQUEST_STATES.CREATED,
      createdAt: '2020-10-07T13:01:11.288349Z',
      updatedAt: null,
    })

  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
    managementApiClient,
  })

  const outgoingAccessRequest = await outgoingAccessRequestStore.create({
    organizationName: 'organization-name',
    serviceName: 'service-name',
  })

  expect(
    managementApiClient.managementCreateAccessRequest,
  ).toHaveBeenCalledWith({
    body: {
      organizationName: 'organization-name',
      serviceName: 'service-name',
    },
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

test('updating from server', async () => {
  const outgoingAccessRequestStore = new OutgoingAccessRequestStore({})

  expect(outgoingAccessRequestStore.outgoingAccessRequests.size).toEqual(0)
  expect(outgoingAccessRequestStore.updateFromServer()).toBeNull()

  let outgoingAccessRequestModel = await outgoingAccessRequestStore.updateFromServer(
    {
      id: '42',
      state: ACCESS_REQUEST_STATES.CREATED,
    },
  )

  // new model should be created
  expect(outgoingAccessRequestStore.outgoingAccessRequests.size).toEqual(1)
  expect(outgoingAccessRequestModel).toBeInstanceOf(OutgoingAccessRequestModel)

  outgoingAccessRequestModel = await outgoingAccessRequestStore.updateFromServer(
    {
      id: '42',
      state: ACCESS_REQUEST_STATES.RECEIVED,
    },
  )

  // existing model should be updated
  expect(outgoingAccessRequestStore.outgoingAccessRequests.size).toEqual(1)
  expect(outgoingAccessRequestModel).toBeInstanceOf(OutgoingAccessRequestModel)
  expect(outgoingAccessRequestModel.state).toEqual(
    ACCESS_REQUEST_STATES.RECEIVED,
  )
})
