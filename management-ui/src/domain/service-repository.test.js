// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import ServiceRepository from './service-repository'

describe('the ServiceRepository', () => {
  describe('getting all services', () => {
    afterEach(() => global.fetch.mockRestore())

    it('should return the services', async () => {
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

      const result = await ServiceRepository.getAll()

      expect(result).toEqual([
        {
          name: 'A Service',
        },
      ])

      expect(global.fetch).toHaveBeenCalledWith('/api/v1/services')
    })

    it('should return an empty list when the response is an empty object', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () => Promise.resolve({}),
      })

      const result = await ServiceRepository.getAll()

      expect(result).toEqual([])

      expect(global.fetch).toHaveBeenCalledWith('/api/v1/services')
    })
  })
})
