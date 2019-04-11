import React from 'react'
import {shallow} from 'enzyme'
import { LoginPageContainer } from './index'

describe('LoginPageContainer', () => {
  describe('on initialization', () => {
    describe('when the organization has been loaded', () => {
      it('should fetch the login information', () => {
        const props = {
          fetchIrmaLoginInformation: jest.fn(),
          organization: {
            name: 'foo',
            insight_irma_endpoint: 'irma_endpoint',
            insight_log_endpoint: 'log_endpoint'
          }
        }

        const wrapper = shallow(<LoginPageContainer {...props} />)
        const instance = wrapper.instance()
        expect(instance.props.fetchIrmaLoginInformation).toHaveBeenCalledWith({
          insight_irma_endpoint: 'irma_endpoint',
          insight_log_endpoint: 'log_endpoint'
        })
      })
    })
  })

  describe('when the organization changes', () => {
    it('should re-fetch the login information', () => {
      const props = {
        fetchIrmaLoginInformation: jest.fn(),
        organization: {
          name: 'foo',
          insight_irma_endpoint: 'foo_irma_endpoint',
          insight_log_endpoint: 'foo_log_endpoint'
        }
      }

      const newOrganization = {
        name: 'bar',
        insight_irma_endpoint: 'bar_irma_endpoint',
        insight_log_endpoint: 'bar_log_endpoint'
      }

      const wrapper = shallow(<LoginPageContainer {...props} />)
      const instance = wrapper.instance()
      wrapper.setProps({organization: newOrganization})

      expect(instance.props.fetchIrmaLoginInformation).toHaveBeenNthCalledWith(1, {
        insight_irma_endpoint: 'foo_irma_endpoint',
        insight_log_endpoint: 'foo_log_endpoint'
      })

      expect(instance.props.fetchIrmaLoginInformation).toHaveBeenNthCalledWith(2, {
        insight_irma_endpoint: 'bar_irma_endpoint',
        insight_log_endpoint: 'bar_log_endpoint'
      })
    })
  })
})
