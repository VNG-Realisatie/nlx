import React from 'react'
import { shallow } from 'enzyme'
import LogsPage from './index'
import ErrorMessage from '../ErrorMessage'
import { StyledLogsPage } from './index.styles'

describe('LogsPage', () => {
  let wrapper

  beforeEach(() => {
    wrapper = shallow(<LogsPage organizationName="dummy-name"
                                logs={[{
                                  subjects: [ "foo", "bar" ],
                                  requestedBy: "requestedBy",
                                  requestedAt: "requestedAt",
                                  reason: "reason",
                                  date: new Date()
                                }]}/>)
  })

  it('should show the LogsTable', () => {
    expect(wrapper.is(StyledLogsPage)).toEqual(true)
  })

  describe('when no logs are available', () => {
    it('should show an ErrorMessage', () => {
      wrapper.setProps({ logs: [] })
      expect(wrapper.is(ErrorMessage)).toEqual(true)
    })
  })
})
