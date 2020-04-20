// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import InwayRepository from './inway-repository'

describe('the InwayRepository', () => {
  describe('getting all inways', () => {
    afterEach(() => global.fetch.mockRestore())

    it('should return the inways', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            inways: [
              {
                name: 'An Inway',
                hostname: 'hostname',
                selfAddress: 'self-address',
                version: 'version',
              },
            ],
          }),
      })

      const result = await InwayRepository.getAll()

      expect(result).toEqual([
        {
          name: 'An Inway',
          hostname: 'hostname',
          selfAddress: 'self-address',
          version: 'version',
        },
      ])

      expect(global.fetch).toHaveBeenCalledWith('/api/v1/inways')
    })

    it('should return an empty list when the response is an empty object', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () => Promise.resolve({}),
      })

      const result = await InwayRepository.getAll()

      expect(result).toEqual([])

      expect(global.fetch).toHaveBeenCalledWith('/api/v1/inways')
    })
  })
})
