// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import {
  reduceInwayStatesToStatus,
  mapListServicesAPIResponse,
} from './map-list-services-api-response'

describe('reduce the inway states to a service status', () => {
  describe('when there are no inways', () => {
    it('should return down', () => {
      const inways = []
      const status = reduceInwayStatesToStatus(inways)
      expect(status).toBe('down')
    })
  })

  describe('when there is only one inway', () => {
    it('should return the state of the inway', () => {
      const inways = [
        {
          state: 'STATE_UP',
        },
      ]
      const status = reduceInwayStatesToStatus(inways)
      expect(status).toBe('up')
    })
  })

  describe('when there are multiple inways with the same state', () => {
    it('should return the same state', () => {
      const inways = [
        {
          state: 'STATE_UP',
        },
        {
          state: 'STATE_UP',
        },
      ]

      const status = reduceInwayStatesToStatus(inways)
      expect(status).toBe('up')
    })
  })

  describe('when there are two inways with a different state', () => {
    it('should return degraded', () => {
      const inways = [
        {
          state: 'STATE_UP',
        },
        {
          state: 'STATE_DOWN',
        },
      ]
      const status = reduceInwayStatesToStatus(inways)
      expect(status).toBe('degraded')
    })
  })

  describe('when there is an inway with a state that is not known', () => {
    it('should return unknown', () => {
      const inways = [
        {
          state: 'STATE_UNSPECIFIED',
        },
      ]
      const status = reduceInwayStatesToStatus(inways)
      expect(status).toBe('unknown')
    })
  })
})

describe('mapping the API response', () => {
  it('should map the properties without costs', () => {
    const apiResponse = {
      services: [
        {
          /* eslint-disable camelcase */
          organization: {
            name: 'foo',
            serial_number: '00000000000000000000',
          },
          name: 'bar',
          inway_addresses: ['https://www.duck.com'],
          documentation_url: 'https://www.duck.com',
          api_specification_type: 'openapi',
          public_support_contact: 'foo@bar.baz',
          /* eslint-enable camelcase */
        },
      ],
    }

    expect(mapListServicesAPIResponse(apiResponse)).toEqual([
      {
        organization: {
          name: 'foo',
          serialNumber: '00000000000000000000',
        },
        name: 'bar',
        apiType: 'openapi',
        contactEmailAddress: 'foo@bar.baz',
        documentationUrl: 'https://www.duck.com',
        status: 'down',
        oneTimeCosts: 0,
        monthlyCosts: 0,
        requestCosts: 0,
      },
    ])
  })

  it('should map the properties with costs', () => {
    const apiResponse = {
      services: [
        {
          /* eslint-disable camelcase */
          organization: {
            name: 'foo',
            serial_number: '00000000000000000000',
          },
          name: 'bar',
          inway_addresses: ['https://www.duck.com'],
          documentation_url: 'https://www.duck.com',
          api_specification_type: 'openapi',
          public_support_contact: 'foo@bar.baz',
          costs: {
            one_time: 1,
            monthly: 2,
            request: 3,
          },
          /* eslint-enable camelcase */
        },
      ],
    }

    expect(mapListServicesAPIResponse(apiResponse)).toEqual([
      {
        organization: {
          name: 'foo',
          serialNumber: '00000000000000000000',
        },
        name: 'bar',
        apiType: 'openapi',
        contactEmailAddress: 'foo@bar.baz',
        documentationUrl: 'https://www.duck.com',
        status: 'down',
        oneTimeCosts: 0.01,
        monthlyCosts: 0.02,
        requestCosts: 0.03,
      },
    ])
  })

  describe('when the response does not contain the services object', () => {
    it('should return an empty array', () => {
      const apiResponse = {}
      expect(mapListServicesAPIResponse(apiResponse)).toEqual([])
    })
  })
})
