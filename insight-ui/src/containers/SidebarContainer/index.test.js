// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import Sidebar from '../../components/Sidebar'
import { SidebarContainer } from './index'

describe('SidebarContainer', () => {
  describe('on initialization', () => {
    let wrapper
    let instance

    beforeEach(() => {
      const props = {
        fetchOrganizationsRequest: jest.fn(),
      }

      wrapper = shallow(<SidebarContainer {...props} />)
      instance = wrapper.instance()
    })

    it('should have an empty query value', () => {
      expect(wrapper.state('query')).toEqual('')
    })

    it('should fetch the organizations', () => {
      expect(instance.props.fetchOrganizationsRequest).toHaveBeenCalled()
    })
  })

  describe('filtering the organizations by query', () => {
    let instance

    beforeEach(() => {
      const wrapper = shallow(<SidebarContainer />)
      instance = wrapper.instance()
    })

    it('should include the name of the organizations', () => {
      const organizations = [{ name: 'foo' }]
      expect(
        instance.getFilteredOrganizationsByQuery(organizations, ''),
      ).toEqual(['foo'])
    })

    it('should filter by organization name', () => {
      const organizations = [{ name: 'foo' }, { name: 'bar' }, { name: 'baz' }]
      expect(
        instance.getFilteredOrganizationsByQuery(organizations, 'ba'),
      ).toEqual(['bar', 'baz'])
    })

    it('should be case-insensitive', () => {
      const organizations = [{ name: 'foo' }, { name: 'bar' }, { name: 'baz' }]
      expect(
        instance.getFilteredOrganizationsByQuery(organizations, 'FoO'),
      ).toEqual(['foo'])
    })
  })

  describe('the organizations shown in the sidebar', () => {
    it('should be the filtered organizations by query', () => {
      jest
        .spyOn(SidebarContainer.prototype, 'getFilteredOrganizationsByQuery')
        .mockImplementation(() => ['filteredByQuery'])

      const wrapper = shallow(<SidebarContainer />)
      expect(
        wrapper.instance().getFilteredOrganizationsByQuery,
      ).toHaveBeenCalled()
      expect(wrapper.find(Sidebar).prop('organizations')).toEqual([
        'filteredByQuery',
      ])
    })
  })
})
