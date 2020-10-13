// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessGrantRepository from './access-grant-repository'

describe('getting access grants per service', () => {
  afterEach(() => global.fetch.mockRestore())

  it('should return a list of grants', async () => {
    const accessGrants = [
      {
        id: '1234abcd',
        serviceName: 'service',
        organizationName: 'Organization',
        publicKeyFingerprint: 'printFinger=',
        createdAt: '2020-10-07T13:01:11.288349Z',
      },
    ]

    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      status: 200,
      json: () => Promise.resolve({ accessGrants }),
    })

    const result = await AccessGrantRepository.getByServiceName('service')

    expect(result).toEqual(accessGrants)
    expect(global.fetch).toHaveBeenCalledWith(
      '/api/v1/access-grants/services/service',
    )
  })

  it('should return an empty list if there are no access grants', async () => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      status: 200,
      json: () => Promise.resolve({}),
    })

    const result = await AccessGrantRepository.getByServiceName('service')

    expect(result).toEqual([])
  })
})

describe('revoking an access grant', () => {
  afterEach(() => global.fetch.mockRestore())

  it('should return null', async () => {
    jest.spyOn(global, 'fetch').mockResolvedValue({
      ok: true,
      status: 200,
    })

    const result = await AccessGrantRepository.revokeAccessGrant({
      organizationName: 'organization-name',
      serviceName: 'service-name',
      accessGrantId: 'access-grant-id',
    })

    expect(global.fetch).toHaveBeenCalledWith(
      '/api/v1/access-grants/service/service-name/organizations/organization-name/access-grant-id/revoke',
      {
        method: 'POST',
        body: JSON.stringify({
          organizationName: 'organization-name',
          serviceName: 'service-name',
          accessGrantID: 'access-grant-id',
        }),
      },
    )

    expect(result).toBeNull()
  })
})
