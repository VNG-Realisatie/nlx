import React from 'react'
import {shallow} from 'enzyme'
import { SidebarContainer } from './index'

describe('SidebarContainer', () => {
  describe('on initialization', () => {
    it('should fetch the organizations', () => {
      const props = {
        fetchOrganizationsRequest: jest.fn()
      }
      const wrapper = shallow(<SidebarContainer {...props} />)
      const instance = wrapper.instance()
      expect(instance.props.fetchOrganizationsRequest).toHaveBeenCalled()
    })
  })
})
