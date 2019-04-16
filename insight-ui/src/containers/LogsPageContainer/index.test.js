import React from 'react'
import {shallow} from 'enzyme'
import { LogsPageContainer } from './index'

describe('LogsPageContainer', () => {
  describe('on initialization', () => {
    let wrapper
    let instance

    beforeEach(() => {
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
      instance = wrapper.instance()
    })

    it('should have an empty query value', () => {
      expect(wrapper.state('query')).toEqual('')
    })

    it('should fetch the organization logs', () => {
      expect(instance.props.fetchOrganizationLogs).toHaveBeenCalledWith({
        proofUrl: 'proof_url',
        insight_log_endpoint: 'log_endpoint'
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

  describe('filtering the logs by query', () => {
    let wrapper
    let instance

    beforeEach(() => {
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

      wrapper = shallow(<LogsPageContainer {...props} />)
      instance = wrapper.instance()
    })

    it('should filter case-insensitive by the subjects property', () => {
      const logs = [ { subjects: ['foo', 'bar'] }, { subjects: ['baz', 'foo'] } ]
      expect(instance.getFilteredLogsByQuery(logs, 'bAr')).toEqual([
        { subjects: ['foo', 'bar'] }
      ])
    })

    it('should filter case-insensitive by the requestedBy property', () => {
      const logs = [ { requestedBy: 'foo'}, { requestedBy: 'bar'}, { requestedBy: 'baz'} ]
      expect(instance.getFilteredLogsByQuery(logs, 'BA')).toEqual([
        { requestedBy: 'bar' },
        { requestedBy: 'baz' }
      ])
    })

    it('should filter case-insensitive by the requestedBy property', () => {
      const logs = [ { requestedAt: 'foo'}, { requestedAt: 'bar'}, { requestedAt: 'baz'} ]
      expect(instance.getFilteredLogsByQuery(logs, 'BA')).toEqual([
        { requestedAt: 'bar' },
        { requestedAt: 'baz' }
      ])
    })

    it('should filter case-insensitive by the reason', () => {
      const logs = [ { reason: 'foo'}, { reason: 'bar'}, { reason: 'baz'} ]
      expect(instance.getFilteredLogsByQuery(logs, 'BA')).toEqual([
        { reason: 'bar' },
        { reason: 'baz' }
      ])
    })
  })
})
