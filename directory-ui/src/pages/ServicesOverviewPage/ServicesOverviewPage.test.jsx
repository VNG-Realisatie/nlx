import React from 'react'
import { shallow } from 'enzyme'
import { Spinner } from '@commonground/design-system'
import ServicesOverviewPage from './ServicesOverviewPage';
import ErrorMessage from '../../components/ErrorMessage/ErrorMessage';
import ServicesTableContainer from '../../containers/ServicesTableContainer/ServicesTableContainer'

describe('ServicesOverviewPage', () => {
  let wrapper
  let instance

  beforeEach(() => {
    wrapper = shallow(<ServicesOverviewPage/>)
    instance = wrapper.instance()
  })

  describe('the initial state', () => {
    it('should have no filters applied', () => {
      expect(wrapper.state()).toMatchObject({
        query: '',
        displayOfflineServices: true
      })
    })
  })

  describe('loading the services', () => {
    it('should show a spinner', () => {
      wrapper.setState({
        loading: true
      })

      expect(wrapper.contains(<Spinner />)).toBe(true)
    })
  })

  describe('failing to load the services', () => {
    it('should show an error message', () => {
      wrapper.setState({
        loading: false,
        error: true
      })

      expect(wrapper.contains(<ErrorMessage />)).toBe(true)
    })
  })

  describe('when the services are loaded', () => {
    it('should show the services table', () => {
      wrapper.setState({
        loading: false,
        error: false,
        services: []
      })

      expect(wrapper.exists('ServicesTableContainer')).toBe(true)
    })
  })

  describe('the escape button', () => {
    it('should clear the search query', () => {
      const ESCAPE_KEY_CODE = 27

      wrapper.setState({ query: 'foo' })

      global.document.dispatchEvent(new KeyboardEvent('keydown', {
        keyCode: ESCAPE_KEY_CODE
      }))

      expect(wrapper.state('query')).toBe('')
    })
  })
})
