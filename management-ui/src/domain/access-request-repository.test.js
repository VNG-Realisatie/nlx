// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessRequestRepository from './access-request-repository'

describe('creating an access request', () => {
  beforeEach(() => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      status: 201,
      json: () =>
        Promise.resolve({
          id: 'string',
          organizationName: 'organization',
          serviceName: 'service',
          state: 'CREATED',
          createdAt: '2020-06-30T10:36:57.100Z',
          updatedAt: '2020-06-30T10:36:57.100Z',
        }),
    })
  })

  afterEach(() => global.fetch.mockRestore())

  it('should return the services', async () => {
    const result = await AccessRequestRepository.createAccessRequest({
      organizationName: 'organization',
      serviceName: 'service',
    })

    expect(result).toEqual({
      id: 'string',
      organizationName: 'organization',
      serviceName: 'service',
      state: 'CREATED',
      createdAt: '2020-06-30T10:36:57.100Z',
      updatedAt: '2020-06-30T10:36:57.100Z',
    })
  })

  it('rejects duplicate requests', async () => {
    jest.spyOn(global, 'fetch').mockImplementation(async () => ({
      ok: false,
      status: 409,
    }))

    await expect(
      AccessRequestRepository.createAccessRequest({
        organizationName: 'organization',
        serviceName: 'service',
      }),
    ).rejects.toThrowError(
      /^Request already sent, please refresh the page to see the latest state\.$/,
    )
  })
})

describe('list incoming access requests', () => {
  it('should return a list of access requests', async () => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      status: 200,
      json: () => Promise.resolve([]),
    })

    const result = await AccessRequestRepository.listIncomingAccessRequests(
      'service-name',
    )

    expect(global.fetch).toHaveBeenCalledWith(
      '/api/v1/access-requests/incoming/services/service-name',
    )

    expect(result).toEqual([])
  })
})

describe('approve an incoming access requests', () => {
  it('should return an empty promise', async () => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      status: 200,
      json: () => Promise.resolve([]),
    })

    const result = await AccessRequestRepository.approveIncomingAccessRequest({
      serviceName: 'service-name',
      id: '42',
    })

    expect(global.fetch).toHaveBeenCalledWith(
      '/api/v1/access-requests/incoming/services/service-name/42/approve',
      expect.objectContaining({
        method: 'POST',
      }),
    )

    expect(result).toBeNull()
  })
})

describe('sending an access request', () => {
  it('should return an empty promise', async () => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      status: 200,
      json: async () => null,
    })

    const parameters = {
      organizationName: 'organization-name',
      serviceName: 'service-name',
      id: 'access-request-id',
    }

    const result = await AccessRequestRepository.sendAccessRequest(parameters)

    expect(global.fetch).toHaveBeenCalledWith(
      '/api/v1/access-requests/outgoing/organizations/organization-name/services/service-name/access-request-id/send',
      expect.objectContaining({
        method: 'POST',
      }),
    )

    expect(result).toBeNull()
  })
})
