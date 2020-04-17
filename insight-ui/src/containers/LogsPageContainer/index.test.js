// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'
import { LogsPageContainer, getPageFromQueryString } from './index'

describe('LogsPageContainer', () => {
  describe('on initialization', () => {
    let wrapper
    let instance

    beforeEach(() => {
      const props = {
        fetchOrganizationLogs: jest.fn(),
        organization: {
          name: 'foo',
          insight_irma_endpoint: 'irma_endpoint', // eslint-disable-line camelcase
          insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
        },
        proof: 'the_proof',
      }

      wrapper = shallow(<LogsPageContainer {...props} />)
      instance = wrapper.instance()
    })

    it('should fetch the organization logs', () => {
      expect(instance.props.fetchOrganizationLogs).toHaveBeenCalledWith({
        proof: 'the_proof',
        insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
        page: 0,
      })
    })
  })

  describe('when the organization changes', () => {
    it('should re-fetch the login information', () => {
      const props = {
        fetchOrganizationLogs: jest.fn(),
        organization: {
          name: 'foo',
          insight_irma_endpoint: 'foo_irma_endpoint', // eslint-disable-line camelcase
          insight_log_endpoint: 'foo_log_endpoint', // eslint-disable-line camelcase
        },
        proof: 'the_proof',
      }

      const newOrganization = {
        name: 'bar',
        insight_irma_endpoint: 'bar_irma_endpoint', // eslint-disable-line camelcase
        insight_log_endpoint: 'bar_log_endpoint', // eslint-disable-line camelcase
      }

      const wrapper = shallow(<LogsPageContainer {...props} />)
      const instance = wrapper.instance()
      wrapper.setProps({ organization: newOrganization })

      expect(instance.props.fetchOrganizationLogs).toHaveBeenNthCalledWith(1, {
        proof: 'the_proof',
        insight_log_endpoint: 'foo_log_endpoint', // eslint-disable-line camelcase
        page: 0,
      })

      expect(instance.props.fetchOrganizationLogs).toHaveBeenNthCalledWith(2, {
        proof: 'the_proof',
        insight_log_endpoint: 'bar_log_endpoint', // eslint-disable-line camelcase
        page: 0,
      })
    })
  })
})

describe('get page from the query string', () => {
  it.each([
    ['', undefined],
    ['?page=0', 0],
    ['?page=1', 1],
  ])('%s', (queryString, expected) => {
    const result = getPageFromQueryString(queryString)
    expect(result).toEqual(expected)
  })
})
