// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import fetchMock from 'jest-fetch-mock'
import { clear } from './async-memoize'
import { fetchWithCaching, fetchWithoutCaching } from './fetch-utils'

export const resetFetchWithCaching = () => clear(fetchWithCaching.memo)

afterEach(() => global.fetch.mockRestore())

test('fetchWithoutCaching should use headers to prevent caching', async () => {
  jest.spyOn(global, 'fetch').mockResolvedValue({
    ok: true,
    status: 200,
    json: async () => ({ data: 'value' }),
  })

  const result = await fetchWithoutCaching('/dynamic-resource')

  expect(await result.json()).toEqual({ data: 'value' })

  expect(global.fetch).toHaveBeenCalledWith('/dynamic-resource', {
    headers: {
      'Cache-Control': 'no-cache',
      Pragma: 'no-cache',
      Expires: 'Sat, 01 Jan 2000 00:00:00 GMT',
    },
  })
})

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
