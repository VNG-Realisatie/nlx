// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import ServicesTableRow, { apiUrlForService } from './ServicesTableRow'

test('renders without crashing', () => {
  expect(() =>
    render(
      <table>
        <tbody>
          <ServicesTableRow
            name="service"
            organization="organization"
            status="unknown"
          />
        </tbody>
      </table>,
    ),
  ).not.toThrow()
})

describe('the API address', () => {
  it('should consist out of the organization and service name', () => {
    const apiAddress = apiUrlForService('organization', 'service')
    expect(apiAddress).toBe('http://{your-outway-address}/organization/service')
  })
})
