// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import ServiceRepository from './service-repository'

describe('the ServiceRepository', () => {
  describe('updating a service', () => {
    describe('when the payload is correct', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(async () => ({
          ok: true,
          status: 200,
          json: async () => null,
        }))
      })

      afterEach(() => global.fetch.mockRestore())
      it('should return successfully', async () => {
        const result = await ServiceRepository.update('my-service', {
          name: 'my-service',
        })
        await expect(result).toBeNull()
        expect(global.fetch).toHaveBeenCalledWith(
          '/api/v1/services/my-service',
          {
            method: 'PUT',
            body: JSON.stringify({ name: 'my-service' }),
          },
        )
      })
    })
  })
})
