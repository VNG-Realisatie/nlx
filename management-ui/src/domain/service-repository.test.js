// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import ServiceRepository from './service-repository'

describe('the ServiceRepository', () => {
  describe('creating a service', () => {
    describe('when the creation is successful', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: true,
            status: 201,
            json: () => Promise.resolve({}),
          }),
        )
      })

      afterEach(() => global.fetch.mockRestore())

      it('should return an empty object', async () => {
        const result = await ServiceRepository.create({ name: 'my-service' })
        await expect(result).toEqual({})
        expect(global.fetch).toHaveBeenCalledWith('/api/v1/services', {
          method: 'POST',
          body: JSON.stringify({ name: 'my-service' }),
        })
      })
    })

    describe('with invalid user input', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: false,
            status: 400,
          }),
        )
      })

      afterEach(() => global.fetch.mockRestore())

      it('should throw an error', async () => {
        const create = ServiceRepository.create('invalid argument')
        await expect(create).rejects.toEqual(new Error('invalid user input'))

        expect(global.fetch).toHaveBeenCalledWith('/api/v1/services', {
          method: 'POST',
          body: '"invalid argument"',
        })
      })
    })

    describe('when an unexpected error happens', () => {
      beforeEach(() => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: false,
          }),
        )
      })

      afterEach(() => global.fetch.mockRestore())

      it('should throw an error', async () => {
        const create = ServiceRepository.create()

        await expect(create).rejects.toEqual(
          new Error('unable to handle the request'),
        )

        expect(global.fetch).toHaveBeenCalledWith('/api/v1/services', {
          method: 'POST',
          body: undefined,
        })
      })
    })
  })

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
