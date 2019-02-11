import React from 'react'
import { shallow } from 'enzyme'
import ServicesTable from './ServicesTable'

describe('ServicesTable', () => {
  let wrapper
  let instance

  beforeEach(() => {
    wrapper = shallow(<ServicesTable/>)
    instance = wrapper.instance()
  })

  describe('sorting', () => {
    it('should have null values for default sorting', () => {
      expect(wrapper.state('sortBy')).toBeNull()
      expect(wrapper.state('sortOrder')).toBeNull()
    })
  })

  describe('toggle the sorting', () => {
    describe('when no sorting was active before', () => {
      it('should sort ascending', () => {
        instance.toggleSorting('organization')
        expect(wrapper.state('sortBy')).toBe('organization')
        expect(wrapper.state('sortOrder')).toBe('asc')
      })
    })

    describe('when the column was already sorted on', () => {
      it('should reverse the sorting', () => {
        instance.toggleSorting('organization')
        expect(wrapper.state('sortOrder')).toBe('asc')

        instance.toggleSorting('organization')
        expect(wrapper.state('sortOrder')).toBe('desc')

        instance.toggleSorting('organization')
        expect(wrapper.state('sortOrder')).toBe('asc')
      })
    })
  })
})
