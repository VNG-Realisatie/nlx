import React from 'react'
import { shallow } from 'enzyme'
import { Spinner } from '@commonground/design-system'
import ServicesOverviewPage, { mapListServicesAPIResponse } from "./ServicesOverviewPage";
import ErrorMessage from "../../components/ErrorMessage/ErrorMessage";
import ServicesTableContainer from "../../containers/ServicesTableContainer/ServicesTableContainer";

describe('ServicesOverviewPage', () => {
  let wrapper
  let instance

  beforeEach(() => {
    wrapper = shallow(<ServicesOverviewPage/>)
    instance = wrapper.instance()
  })

  describe('mapping the API response', () => {
    it('should map the properties', () => {
      const apiResponse = {
        services: [{
          organization_name: 'foo',
          service_name: 'bar',
          inway_addresses: ['https://www.duck.com'],
          documentation_url: 'https://www.duck.com',
          api_specification_type: 'openapi',
        }]
      }

      expect(mapListServicesAPIResponse(apiResponse)).toEqual([{
        organization: 'foo',
        name: 'bar',
        status: 'online',
        documentationLink: 'https://www.duck.com',
        apiType: 'openapi'
      }])
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

      expect(wrapper.contains(<ServicesTableContainer />)).toBe(true)
    })
  })

  describe('the escape button', () => {
    it('should clear the search query', () => {
      const ESCAPE_KEY_CODE = 27

      wrapper.setState({
        displayOnlyContaining: 'foo'
      })

      global.document.dispatchEvent(new KeyboardEvent('keydown', {
        keyCode: ESCAPE_KEY_CODE
      }))

      expect(wrapper.state('displayOnlyContaining')).toBe('')
    })
  })
})
