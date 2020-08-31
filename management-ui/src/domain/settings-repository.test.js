// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import SettingsRepository from './settings-repository'
import { PREVENT_CACHING_HEADERS } from './fetch-utils'

describe('the SettingsRepository', () => {
  afterEach(() => global.fetch.mockRestore())

  describe('getting the settings', () => {
    it('should return the settings', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            organizationInway: 'inway-name',
          }),
      })

      const result = await SettingsRepository.get()

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/settings',
        expect.objectContaining({
          headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
        }),
      )

      expect(result).toEqual({
        organizationInway: 'inway-name',
      })
    })

    describe('when an unexpected error happens', () => {
      it('should throw an error', async () => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: false,
            status: 500,
          }),
        )

        await expect(SettingsRepository.get()).rejects.toEqual(
          new Error('unable to handle the request'),
        )

        expect(global.fetch).toHaveBeenCalledWith(
          '/api/v1/settings',
          expect.objectContaining({
            headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
          }),
        )
      })
    })
  })

  describe('updating the settings', () => {
    it('should return an empty promise', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: async () => null,
      })

      const updatedSettings = {
        organizationInway: 'another-inway-name',
      }

      const result = await SettingsRepository.update(updatedSettings)

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/settings',
        expect.objectContaining({
          headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
          method: 'PUT',
        }),
      )

      expect(result).toBeNull()
    })

    describe('when an unexpected error happens', () => {
      it('should throw an error', async () => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: false,
            status: 500,
          }),
        )

        await expect(
          SettingsRepository.update({
            organizationInway: '',
          }),
        ).rejects.toEqual(new Error('unable to handle the request'))

        expect(global.fetch).toHaveBeenCalledWith(
          '/api/v1/settings',
          expect.objectContaining({
            headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
          }),
        )
      })
    })
  })
})
