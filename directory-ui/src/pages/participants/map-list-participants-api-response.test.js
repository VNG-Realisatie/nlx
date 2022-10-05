// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import { mapListParticipantsAPIResponse } from './map-list-participants-api-response'

describe('mapping the API response', () => {
  it('should map the properties without costs', () => {
    const apiResponse = {
      participants: [
        {
          /* eslint-disable camelcase */
          organization: {
            name: 'foo',
            serial_number: '00000000000000000000',
          },
          created_at: '2021-01-01T00:00:00',
          statistics: {
            inways: 10,
            outways: 3,
            services: 42,
          },
          /* eslint-enable camelcase */
        },
      ],
    }

    expect(mapListParticipantsAPIResponse(apiResponse)).toEqual([
      {
        organization: {
          name: 'foo',
          serialNumber: '00000000000000000000',
        },
        createdAt: new Date('2021-01-01T00:00:00'),
        inwayCount: 10,
        outwayCount: 3,
        serviceCount: 42,
      },
    ])
  })

  describe('when the response does not contain the participants object', () => {
    it('should return an empty array', () => {
      const apiResponse = {}
      expect(mapListParticipantsAPIResponse(apiResponse)).toEqual([])
    })
  })
})
