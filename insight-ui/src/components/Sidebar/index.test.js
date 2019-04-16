import React from 'react'
import { NavLink } from 'react-router-dom'
import { shallow } from 'enzyme'
import Sidebar from './index'

describe('Sidebar', () => {
  let wrapper
  beforeEach(() => {
    wrapper = shallow(<Sidebar organizations={['a', 'b', 'c']} />)
  })

  it('should render the organizations', () => {
    expect(wrapper.find('li').length).toEqual(3)
  })

  describe('the organizations', () => {
    it('should link to the organization detail page', () => {
      const links = wrapper
        .find('li')
        .map(node => node.find(NavLink).prop('to'))

      expect(links[0]).toEqual('/organization/a')
      expect(links[1]).toEqual('/organization/b')
      expect(links[2]).toEqual('/organization/c')
    })
  })
})
