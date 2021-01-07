// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import fetchMock from 'jest-fetch-mock'
import { clear } from './async-memoize'
import { fetchWithCaching, throwOnError } from './fetch-utils'

export const resetFetchWithCaching = () => clear(fetchWithCaching.memo)

afterEach(() => global.fetch.mockRestore())

describe('fetchMemoized', () => {
  beforeEach(() => {
    resetFetchWithCaching()
    fetchMock.mockResponse(async (request) =>
      JSON.stringify({ url: request.url }),
    )
  })
  afterEach(() => {
    fetchMock.resetMocks()
  })

  it('should call fetch only once for every url', async () => {
    const firstResult = await (await fetchWithCaching('/test')).json()
    expect(firstResult).toEqual({ url: '/test' })
    expect(fetchMock.mock.calls).toHaveLength(1)

    const secondResult = await (await fetchWithCaching('/test')).json()
    expect(firstResult).toEqual(secondResult)
    expect(fetchMock.mock.calls).toHaveLength(1)

    const thirdResult = await (await fetchWithCaching('/other')).json()
    expect(thirdResult).toEqual({ url: '/other' })
    expect(fetchMock.mock.calls).toHaveLength(2)
  })
})

describe('throwOnError', () => {
  it('should pass when everything is fine', () => {
    expect(() =>
      throwOnError({ status: 200, ok: true, json: jest.fn() }),
    ).not.toThrow()
  })

  it('should throw when there is an error', () => {
    expect(() => throwOnError({ status: 404, ok: false })).toThrow()
  })

  it('should throw when the response object is erroneous', () => {
    expect(() => throwOnError()).toThrow()
    expect(() => throwOnError({})).toThrow()
    expect(() => throwOnError({ ok: false })).toThrow()
  })
})
