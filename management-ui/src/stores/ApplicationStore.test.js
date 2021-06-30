// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { ManagementApi } from '../api'
import ApplicationStore, { AUTH_OIDC } from './ApplicationStore'

test('initializing the store', () => {
  const accessGrantStore = new ApplicationStore({
    managementApiClient: new ManagementApi(),
  })

  expect(accessGrantStore.isOrganizationInwaySet).toBeNull()
  expect(accessGrantStore.authStrategy).toEqual(AUTH_OIDC)
})

test.concurrent.each([
  [true, true],
  [false, false],
  ['', false],
  ['inway-name', true],
])('updating isOrganizationInwaySet to %s', (a, expected) => {
  const accessGrantStore = new ApplicationStore({
    managementApiClient: new ManagementApi(),
  })

  accessGrantStore.updateOrganizationInway({
    isOrganizationInwaySet: a,
  })

  expect(accessGrantStore.isOrganizationInwaySet).toBe(expected)
})

describe('the general settings', () => {
  describe('retrieving the settings', () => {
    it('should return the settings', async () => {
      const managementApiClient = new ManagementApi()
      managementApiClient.managementGetSettings = jest.fn().mockResolvedValue({
        inway: 'inway-01',
      })
      const applicationStore = new ApplicationStore({
        rootStore: {},
        managementApiClient,
      })

      expect(await applicationStore.getGeneralSettings()).toEqual({
        inway: 'inway-01',
      })
    })

    describe('when an unexpected error happens', () => {
      it('should throw an error', async () => {
        const managementApiClient = new ManagementApi()
        managementApiClient.managementGetSettings = jest
          .fn()
          .mockRejectedValue('arbitrary error')

        const applicationStore = new ApplicationStore({
          rootStore: {},
          managementApiClient,
        })

        try {
          await applicationStore.getGeneralSettings()
        } catch (e) {}

        expect(applicationStore.error).toEqual('arbitrary error')
      })
    })
  })

  describe('updating the settings', () => {
    it('should return an empty array', async () => {
      const managementApiClient = new ManagementApi()
      managementApiClient.managementUpdateSettings = jest
        .fn()
        .mockResolvedValue([])

      const applicationStore = new ApplicationStore({
        rootStore: {},
        managementApiClient,
      })

      const updateGeneral = await applicationStore.updateGeneralSettings({})

      expect(managementApiClient.managementUpdateSettings).toBeCalledTimes(1)
      expect(updateGeneral).toEqual([])
    })

    describe('when an unexpected error happens', () => {
      it('GGG should throw an error', async () => {
        const managementApiClient = new ManagementApi()
        managementApiClient.managementUpdateSettings = jest
          .fn()
          .mockRejectedValue('arbitrary error')

        const applicationStore = new ApplicationStore({
          rootStore: {},
          managementApiClient,
        })

        try {
          await applicationStore.updateGeneralSettings()
        } catch (e) {}

        expect(
          managementApiClient.managementUpdateSettings,
        ).toHaveBeenCalledTimes(1)

        expect(applicationStore.error).toEqual('arbitrary error')
      })
    })
  })
})
