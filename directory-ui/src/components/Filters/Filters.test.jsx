// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { mount } from 'enzyme'
import Filters from "./Filters";

describe('Filters', () => {
  describe('changing the text input value', () => {
    it('should call the onQueryChanged handler with the query', () => {
      const onQueryChangedSpy = jest.fn()
      const wrapper = mount(<Filters onQueryChanged={onQueryChangedSpy} />)

      wrapper.find('[dataTest="query"] input').simulate('change', {target: {value: 'abc'}})
      expect(onQueryChangedSpy).toHaveBeenCalledWith('abc')
    })
  })

  describe('toggling the offline filter', () => {
    it('should call the onStatusFilterChanged handler with the checked state', () => {
      const onStatusFilterChangedSpy = jest.fn()
      const wrapper = mount(<Filters onStatusFilterChanged={onStatusFilterChangedSpy} />)
      wrapper.find('Switch').simulate('change', {target: {checked: false}});
      expect(onStatusFilterChangedSpy).toHaveBeenCalledWith(false)
    })
  })
})
