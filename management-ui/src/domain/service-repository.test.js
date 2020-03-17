// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import ServiceRepository from './service-repository'

describe('the ServiceRepository', () => {
  describe('getting all services', () => {
    beforeEach(() => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            services: [
              {
                name: 'A Service',
              },
            ],
          }),
      })
    })

    afterEach(() => global.fetch.mockRestore())

    it('should return the services', async () => {
      const result = await ServiceRepository.getAll()

      expect(result).toEqual([
        {
          name: 'A Service',
        },
      ])

      expect(global.fetch).toHaveBeenCalledWith('/api/v1/services')
    })
  })
})
