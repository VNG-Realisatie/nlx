// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import SettingsRepository from './settings-repository'

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

      expect(result).toEqual({
        organizationInway: 'inway-name',
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
      expect(result).toBeNull()
    })
  })
})
