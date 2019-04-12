import React from 'react'
import {shallow} from 'enzyme'
import { LogsPageContainer } from './index'

describe('LogsPageContainer', () => {
  describe('on initialization', () => {
    describe('when the organization has been loaded', () => {
      it('should fetch the organization logs', () => {
        const props = {
          fetchOrganizationLogs: jest.fn(),
          organization: {
            name: 'foo',
            insight_irma_endpoint: 'irma_endpoint',
            insight_log_endpoint: 'log_endpoint'
          },
          loginRequestInfo: {
            proofUrl: 'proof_url'
          }
        }

        const wrapper = shallow(<LogsPageContainer {...props} />)
        const instance = wrapper.instance()
        expect(instance.props.fetchOrganizationLogs).toHaveBeenCalledWith({
          proofUrl: 'proof_url',
          insight_log_endpoint: 'log_endpoint'
        })
      })
    })
  })

  describe('when the organization changes', () => {
    it('should re-fetch the login information', () => {
      const props = {
        fetchOrganizationLogs: jest.fn(),
        organization: {
          name: 'foo',
          insight_irma_endpoint: 'foo_irma_endpoint',
          insight_log_endpoint: 'foo_log_endpoint'
        },
        loginRequestInfo: {
          proofUrl: 'proof_url'
        }
      }

      const newOrganization = {
        name: 'bar',
        insight_irma_endpoint: 'bar_irma_endpoint',
        insight_log_endpoint: 'bar_log_endpoint'
      }

      const wrapper = shallow(<LogsPageContainer {...props} />)
      const instance = wrapper.instance()
      wrapper.setProps({organization: newOrganization})

      expect(instance.props.fetchOrganizationLogs).toHaveBeenNthCalledWith(1, {
        proofUrl: 'proof_url',
        insight_log_endpoint: 'foo_log_endpoint'
      })

      expect(instance.props.fetchOrganizationLogs).toHaveBeenNthCalledWith(2, {
        proofUrl: 'proof_url',
        insight_log_endpoint: 'bar_log_endpoint'
      })
    })
  })
})
