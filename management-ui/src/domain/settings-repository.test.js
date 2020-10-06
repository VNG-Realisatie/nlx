// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import SettingsRepository from './settings-repository'
import { PREVENT_CACHING_HEADERS } from './fetch-utils'

describe('the general settings', () => {
  afterEach(() => global.fetch.mockRestore())

  describe('retrieving the settings', () => {
    it('should return the settings', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            organizationInway: 'inway-name',
          }),
      })

      const result = await SettingsRepository.getGeneralSettings()

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

        await expect(SettingsRepository.getGeneralSettings()).rejects.toEqual(
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

      const result = await SettingsRepository.updateGeneralSettings(
        updatedSettings,
      )

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
          SettingsRepository.updateGeneralSettings({
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

describe('the insight settings', () => {
  afterEach(() => global.fetch.mockRestore())

  describe('retrieving the settings', () => {
    it('should return the settings', async () => {
      jest.spyOn(global, 'fetch').mockResolvedValue({
        ok: true,
        status: 200,
        json: () =>
          Promise.resolve({
            irmaServerURL: 'irma-server-url',
            insightAPIURL: 'insight-api-url',
          }),
      })

      const result = await SettingsRepository.getInsightSettings()

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/insight-configuration',
        expect.objectContaining({
          headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
        }),
      )

      expect(result).toEqual({
        irmaServerURL: 'irma-server-url',
        insightAPIURL: 'insight-api-url',
      })
    })

    describe('when the insight configuration has not been set before', () => {
      it('should return empty values', async () => {
        jest.spyOn(global, 'fetch').mockImplementation(() =>
          Promise.resolve({
            ok: false,
            status: 404,
          }),
        )

        const result = await SettingsRepository.getInsightSettings()

        expect(global.fetch).toHaveBeenCalledWith(
          '/api/v1/insight-configuration',
          expect.objectContaining({
            headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
          }),
        )

        expect(result).toEqual({
          irmaServerURL: '',
          insightAPIURL: '',
        })
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

        await expect(SettingsRepository.getInsightSettings()).rejects.toEqual(
          new Error('unable to handle the request'),
        )

        expect(global.fetch).toHaveBeenCalledWith(
          '/api/v1/insight-configuration',
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
        irmaServerURL: 'irma-server-url',
        insightAPIURL: 'insight-api-url',
      }

      const result = await SettingsRepository.updateInsightSettings(
        updatedSettings,
      )

      expect(global.fetch).toHaveBeenCalledWith(
        '/api/v1/insight-configuration',
        expect.objectContaining({
          headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
          method: 'PUT',
          body:
            '{"irmaServerURL":"irma-server-url","insightAPIURL":"insight-api-url"}',
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
          SettingsRepository.updateInsightSettings({
            irmaServerURL: 'irma-server-url',
            insightAPIURL: 'insight-api-url',
          }),
        ).rejects.toEqual(new Error('unable to handle the request'))

        expect(global.fetch).toHaveBeenCalledWith(
          '/api/v1/insight-configuration',
          expect.objectContaining({
            headers: expect.objectContaining(PREVENT_CACHING_HEADERS),
            body:
              '{"irmaServerURL":"irma-server-url","insightAPIURL":"insight-api-url"}',
          }),
        )
      })
    })
  })
})
