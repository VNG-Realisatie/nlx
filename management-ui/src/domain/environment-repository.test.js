// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import fetchMock from 'jest-fetch-mock'
import EnvironmentRepository from './environment-repository'

describe('the EnvironmentRepository', () => {
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
