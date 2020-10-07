// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import AccessGrantRepository from './access-grant-repository'

describe('the AccessGrantRepository', () => {
  afterEach(() => global.fetch.mockRestore())

  describe('getting access grants per service', () => {
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
})
