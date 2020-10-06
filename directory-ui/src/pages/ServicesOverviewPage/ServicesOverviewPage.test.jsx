// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { shallow } from 'enzyme'

import Spinner from '../../components/Spinner'
import ErrorMessage from '../../components/ErrorMessage/ErrorMessage'
import ServicesOverviewPage from './ServicesOverviewPage'

describe('ServicesOverviewPage', () => {
  let wrapper
  let instance

  beforeEach(() => {
    wrapper = shallow(<ServicesOverviewPage />)
    instance = wrapper.instance()
  })

  describe('the initial state', () => {
    it('should have no filters applied', () => {
      expect(wrapper.state()).toMatchObject({
        query: '',
        displayOfflineServices: true,
      })
    })

    it('should have no service selected', () => {
      expect(wrapper.state()).toMatchObject({
        selectedService: null,
      })
    })
  })

  describe('loading the services', () => {
    it('should show a spinner', () => {
      wrapper.setState({
        loading: true,
      })

      expect(wrapper.contains(<Spinner />)).toBe(true)
    })
  })

  describe('failing to load the services', () => {
    it('should show an error message', () => {
      wrapper.setState({
        loading: false,
        error: true,
      })

      expect(wrapper.contains(<ErrorMessage />)).toBe(true)
    })
  })

  describe('when the services are loaded', () => {
    beforeEach(() => {
      wrapper.setState({
        loading: false,
        error: false,
        services: [],
      })
    })

    it('should show the services table', () => {
      expect(wrapper.exists('ServicesTableContainer')).toBe(true)
    })

    describe('when no service is selected', () => {
      it('should not show the service detail pane', () => {
        wrapper.setState({
          selectedService: null,
        })
        expect(wrapper.exists('ServiceDetailPane')).toBe(false)
      })
    })

    describe('when a service is selected', () => {
      it('should show the service detail pane', () => {
        wrapper.setState({
          selectedService: { foo: 'bar' },
        })
        expect(wrapper.exists('ServiceDetailPane')).toBe(true)
      })
    })
  })

  describe('the escape button', () => {
    it('should clear the search query', () => {
      const ESCAPE_KEY_CODE = 27

      wrapper.setState({ query: 'foo' })

      global.document.dispatchEvent(
        new KeyboardEvent('keydown', {
          keyCode: ESCAPE_KEY_CODE,
        }),
      )

      expect(wrapper.state('query')).toBe('')
    })
  })

  describe('when a query parameter is set', () => {
    beforeEach(() => {
      const mockLocation = { search: '?q=test' }
      wrapper = shallow(<ServicesOverviewPage location={mockLocation} />)
    })

    it('should store the value of the query parameter as internal state', () => {
      expect(wrapper.state('query')).toBe('test')
    })
  })

  describe('the searchOnChangeDebouncable handler', () => {
    let mockHistory
    beforeEach(() => {
      jest.useFakeTimers()

      mockHistory = { push: jest.fn() }
      wrapper = shallow(<ServicesOverviewPage history={mockHistory} />)
      instance = wrapper.instance()
    })

    it('should call a pushState', () => {
      instance.searchOnChangeDebouncable('test')
      jest.runAllTimers()

      expect(mockHistory.push).toHaveBeenCalledWith('?q=test')
    })
  })

  describe('the service clicked handler', () => {
    it('should store the service as local state', () => {
      wrapper.setState({
        selectedService: null,
      })

      wrapper.instance().handleOnServiceClicked({ foo: 'bar' })

      expect(wrapper.state('selectedService')).toEqual({ foo: 'bar' })
    })
  })

  describe('the service detail pane close handler', () => {
    it('should clear the selectedService', () => {
      wrapper.setState({
        selectedService: {},
      })

      wrapper.instance().detailPaneCloseHandler()

      expect(wrapper.state('selectedService')).toBeNull()
    })
  })
})
