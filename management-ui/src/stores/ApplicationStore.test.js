// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { DirectoryServiceApi, ManagementServiceApi } from '../api'
import ApplicationStore, { AUTH_OIDC } from './ApplicationStore'

test('initializing the store', () => {
  const accessGrantStore = new ApplicationStore({
    managementApiClient: new ManagementServiceApi(),
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
    managementApiClient: new ManagementServiceApi(),
  })

  accessGrantStore.updateOrganizationInway({
    isOrganizationInwaySet: a,
  })

  expect(accessGrantStore.isOrganizationInwaySet).toBe(expected)
})

describe('the general settings', () => {
  describe('retrieving the settings', () => {
    it('should return the settings', async () => {
      const managementApiClient = new ManagementServiceApi()
      managementApiClient.managementServiceGetSettings = jest
        .fn()
        .mockResolvedValue({
          settings: {
            inway: 'inway-01',
          },
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
        const managementApiClient = new ManagementServiceApi()
        managementApiClient.managementServiceGetSettings = jest
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
      const managementApiClient = new ManagementServiceApi()
      managementApiClient.managementServiceUpdateSettings = jest
        .fn()
        .mockResolvedValue([])

      const applicationStore = new ApplicationStore({
        rootStore: {},
        managementApiClient,
      })

      const updateGeneral = await applicationStore.updateGeneralSettings({})

      expect(
        managementApiClient.managementServiceUpdateSettings,
      ).toBeCalledTimes(1)
      expect(updateGeneral).toEqual([])
    })

    describe('when an unexpected error happens', () => {
      it('should throw an error', async () => {
        const managementApiClient = new ManagementServiceApi()

        managementApiClient.managementServiceUpdateSettings = jest
          .fn()
          .mockRejectedValue(new Error('arbitrary error'))

        const applicationStore = new ApplicationStore({
          rootStore: {},
          managementApiClient,
        })

        await expect(applicationStore.updateGeneralSettings()).rejects.toThrow(
          'arbitrary error',
        )

        expect(
          managementApiClient.managementServiceUpdateSettings,
        ).toHaveBeenCalledTimes(1)
      })
    })
  })
})

test('the Terms of Service', async () => {
  const directoryApiClient = new DirectoryServiceApi()
  directoryApiClient.directoryServiceGetTermsOfService = jest
    .fn()
    .mockResolvedValue({
      enabled: true,
      url: 'http://example.com',
    })
  const applicationStore = new ApplicationStore({
    rootStore: {},
    directoryApiClient,
  })

  expect(await applicationStore.getTermsOfService()).toEqual({
    enabled: true,
    url: 'http://example.com',
  })
})

test('the Terms of Service status', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceGetTermsOfServiceStatus = jest
    .fn()
    .mockResolvedValue({
      accepted: true,
    })

  const applicationStore = new ApplicationStore({
    rootStore: {},
    managementApiClient,
  })

  expect(await applicationStore.getTermsOfServiceStatus()).toEqual({
    accepted: true,
  })
})

test('accepting the Terms of Service', async () => {
  const managementApiClient = new ManagementServiceApi()
  managementApiClient.managementServiceAcceptTermsOfService = jest
    .fn()
    .mockResolvedValue({})

  const applicationStore = new ApplicationStore({
    rootStore: {},
    managementApiClient,
  })

  expect(await applicationStore.acceptTermsOfService()).toEqual({})
})
