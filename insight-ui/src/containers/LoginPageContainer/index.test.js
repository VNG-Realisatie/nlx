// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import { LoginPageContainer } from './index'

describe('LoginPageContainer', () => {
  describe('on initialization', () => {
    let wrapper
    let instance

    beforeEach(() => {
      const props = {
        fetchIrmaLoginInformation: jest.fn(),
        resetLoginInformation: jest.fn(),
        organization: {
          name: 'foo',
          insight_irma_endpoint: 'irma_endpoint',
          insight_log_endpoint: 'log_endpoint',
        },
      }

      wrapper = shallow(<LoginPageContainer {...props} />)
      instance = wrapper.instance()
    })

    it('should reset the login information', () => {
      expect(instance.props.resetLoginInformation).toHaveBeenCalledTimes(1)
    })

    it('should fetch the login information', () => {
      expect(instance.props.fetchIrmaLoginInformation).toHaveBeenCalledWith({
        insight_irma_endpoint: 'irma_endpoint',
        insight_log_endpoint: 'log_endpoint',
      })
    })
  })

  describe('when the organization changes', () => {
    it('should re-fetch the login information', () => {
      const props = {
        fetchIrmaLoginInformation: jest.fn(),
        resetLoginInformation: () => {},
        organization: {
          name: 'foo',
          insight_irma_endpoint: 'foo_irma_endpoint',
          insight_log_endpoint: 'foo_log_endpoint',
        },
      }

      const newOrganization = {
        name: 'bar',
        insight_irma_endpoint: 'bar_irma_endpoint',
        insight_log_endpoint: 'bar_log_endpoint',
      }

      const wrapper = shallow(<LoginPageContainer {...props} />)
      const instance = wrapper.instance()
      wrapper.setProps({ organization: newOrganization })

      expect(instance.props.fetchIrmaLoginInformation).toHaveBeenNthCalledWith(
        1,
        {
          insight_irma_endpoint: 'foo_irma_endpoint',
          insight_log_endpoint: 'foo_log_endpoint',
        },
      )

      expect(instance.props.fetchIrmaLoginInformation).toHaveBeenNthCalledWith(
        2,
        {
          insight_irma_endpoint: 'bar_irma_endpoint',
          insight_log_endpoint: 'bar_log_endpoint',
        },
      )
    })
  })

  describe('when the login status changes', () => {
    let wrapper
    let instance

    beforeEach(() => {
      const props = {
        history: {
          push: jest.fn(),
        },
        fetchIrmaLoginInformation: () => {},
        resetLoginInformation: () => {},
        organization: {
          name: 'foo',
          insight_irma_endpoint: 'foo_irma_endpoint',
          insight_log_endpoint: 'foo_log_endpoint',
        },
      }

      wrapper = shallow(<LoginPageContainer {...props} />)
      instance = wrapper.instance()
    })

    describe('when the login is successful', () => {
      beforeEach(() => {
        const loginStatus = {
          error: false,
          response: 'DONE',
        }

        wrapper.setProps({ loginStatus })
      })

      describe('the proof is not loaded yet', () => {
        it('should not redirect to the view logs page', () => {
          wrapper.setProps({ proof: { loaded: false } })
          expect(instance.props.history.push).not.toHaveBeenCalled()
        })
      })

      describe('the proof is loaded', () => {
        it('should redirect to the view logs page', () => {
          wrapper.setProps({ proof: { loaded: true } })
          expect(instance.props.history.push).toHaveBeenNthCalledWith(
            1,
            '/organization/foo/logs',
          )
        })
      })
    })
  })
})
