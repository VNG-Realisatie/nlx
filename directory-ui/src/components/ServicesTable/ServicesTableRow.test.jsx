// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'
import copy from 'copy-text-to-clipboard'
import ServicesTableRow, { apiUrlForService } from './ServicesTableRow'
jest.mock('copy-text-to-clipboard')

test('renders without crashing', () => {
  expect(() => {
    shallow(
      <ServicesTableRow
        status="up"
        organization="organization"
        name="service name"
      />,
    )
  }).not.toThrow()
})

describe('the API address', () => {
  it('should consist out of the organization and service name', () => {
    const apiAddress = apiUrlForService('organization', 'service')
    expect(apiAddress).toBe('http://{your-outway-address}/organization/service')
  })
})

describe('clicking the link icon', () => {
  it('should copy the API address to the clipboard', () => {
    const wrapper = shallow(
      <ServicesTableRow
        status="up"
        organization="organization"
        name="service"
      />,
    )
    wrapper.find('[dataTest="link-icon"]').simulate('click')
    expect(copy).toHaveBeenCalledWith(
      'http://{your-outway-address}/organization/service',
    )
  })
})
