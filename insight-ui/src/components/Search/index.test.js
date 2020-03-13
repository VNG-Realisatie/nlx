import React from 'react'
import { mount } from 'enzyme'
import Search from './index'

describe('Search', () => {
  describe('changing the text input value', () => {
    it('should call the onQueryChanged handler with the query', () => {
      const onQueryChangedSpy = jest.fn()
      const wrapper = mount(<Search onQueryChanged={onQueryChangedSpy} />)

      wrapper
        .find('[dataTest="query"] input')
        .simulate('change', { target: { value: 'abc' } })
      expect(onQueryChangedSpy).toHaveBeenCalledWith('abc')
    })
  })
})
