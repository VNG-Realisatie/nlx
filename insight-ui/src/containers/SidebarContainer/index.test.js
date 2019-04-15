import React from 'react'
import {shallow} from 'enzyme'
import { SidebarContainer } from './index'

describe('SidebarContainer', () => {
  describe('on initialization', () => {
    let wrapper
    let instance

    beforeEach(() => {
      const props = {
        fetchOrganizationsRequest: jest.fn()
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

  describe('organizations shown in the sidebar', () => {
    let wrapper
    let instance

    beforeEach(() => {
      const organizations = [
        { name: 'foo'},
        { name: 'bar'},
        { name: 'baz'},
      ]

      wrapper = shallow(<SidebarContainer organizations={organizations} />)
      instance = wrapper.instance()
    })

    it('should include the name of the organizations', () => {
      expect(instance.getOrganizationsForSidebar()).toEqual(['foo', 'bar', 'baz'])
    })

    describe('when the query is set', () => {
      describe('filtering', () => {
        it('should be done by organization name', () => {
          instance.setState({ query: 'ba' })
          expect(instance.getOrganizationsForSidebar()).toEqual(['bar', 'baz'])
        })

        it('should be case-insensitive', () => {
          instance.setState({ query: 'FoO' })
          expect(instance.getOrganizationsForSidebar()).toEqual(['foo'])
        })
      })
    })
  })
})
