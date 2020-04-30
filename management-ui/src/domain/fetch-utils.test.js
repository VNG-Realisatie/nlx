// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { fetchWithoutCaching } from './fetch-utils'

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
