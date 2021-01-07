// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

afterEach(() => global.fetch.mockRestore())

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
