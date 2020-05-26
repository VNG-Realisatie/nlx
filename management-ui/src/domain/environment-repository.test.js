// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import fetchMock from 'jest-fetch-mock'
import EnvironmentRepository from './environment-repository'
import { resetFetchWithCaching } from './fetch-utils.test'

describe('the EnvironmentRepository', () => {
  beforeEach(() => {
    resetFetchWithCaching()
  })
  afterEach(() => {
    fetchMock.resetMocks()
  })

  describe('getting the environment', () => {
    describe('when the api is up', () => {
      beforeEach(() => {
        fetchMock.mockResponses(JSON.stringify({ organizationName: 'test' }))
      })

      it('should return the environment', async () => {
        expect(await EnvironmentRepository.getCurrent()).toEqual({
          organizationName: 'test',
        })
      })

      it('should be called only once', async () => {
        const firstResult = await EnvironmentRepository.getCurrent()
        expect(firstResult).toEqual({ organizationName: 'test' })

        const secondResult = await EnvironmentRepository.getCurrent()
        expect(firstResult).toEqual(secondResult)
        expect(fetchMock.mock.calls).toHaveLength(1)
      })
    })

    describe('when an unexpected error happens', () => {
      it('should throw an error', async () => {
        fetchMock.mockResponse('""', {
          status: 500,
          statusText: 'server error',
        })

        await expect(EnvironmentRepository.getCurrent()).rejects.toEqual(
          new Error('unable to handle the request'),
        )
      })
    })
  })
})
