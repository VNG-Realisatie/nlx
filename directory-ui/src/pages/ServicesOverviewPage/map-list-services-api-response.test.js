// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { mapListServicesAPIResponse } from './map-list-services-api-response'

describe('mapping the API response', () => {
    it('should map the properties', () => {
        const apiResponse = {
            services: [
                {
                    /* eslint-disable camelcase */
                    organization_name: 'foo',
                    service_name: 'bar',
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
                organization: 'foo',
                name: 'bar',
                status: 'online',
                apiType: 'openapi',
                contactEmailAddress: 'foo@bar.baz',
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
