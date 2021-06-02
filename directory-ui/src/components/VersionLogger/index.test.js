// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { render, act } from '@testing-library/react'
import VersionLogger from './index'

describe('VersionLogger', () => {
  beforeEach(() => {
    jest.spyOn(global, 'fetch').mockResolvedValueOnce({
      ok: true,
      status: 200,
      json: () => Promise.resolve({ tag: 'test' }),
    })
  })

  afterEach(() => global.fetch.mockRestore())

  describe('on initialization', () => {
    it('should fetch the version and call the logger with the retrieved tag', async () => {
      const logger = jest.fn()

      // Only way I could get this test to work and not log errors
      await act(async () => {
        await render(<VersionLogger logger={logger} />)
      })

      expect(global.fetch).toHaveBeenCalledTimes(1)
      expect(logger).toHaveBeenCalledWith('test')
    })
  })
})
