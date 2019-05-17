// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import copy from 'copy-text-to-clipboard'
import ServicesTableRow, { apiAddressForService } from './ServicesTableRow'
jest.mock('copy-text-to-clipboard')

it('renders without crashing', () => {
  shallow(<ServicesTableRow status="online"
                            organization="organization"
                            name="service name"
  />)
})

describe('the API address', () => {
  it('should consist out of the organization and service name', () => {
    const apiAddress = apiAddressForService('service', 'organization')
    expect(apiAddress).toBe('http://{your-outway-address}:12018/organization/service')
  })
})

describe('clicking the link icon', () => {
  it('should copy the API address to the clipboard', () => {
    const wrapper = shallow(<ServicesTableRow status="online"
                                              organization="organization"
                                              name="service"
    />)
    wrapper.find('[dataTest="link-icon"]').simulate('click')
    expect(copy).toHaveBeenCalledWith('http://{your-outway-address}:12018/organization/service')
  })
})
