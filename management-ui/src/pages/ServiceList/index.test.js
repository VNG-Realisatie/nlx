// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import { flushPromises } from '../../testHelpers'
import ServiceList from './index'

const mockServiceList = {
    services: [{ name: 'service1' }, { name: 'service2' }],
}

describe('ServiceList', () => {
    describe('the component is initialized', () => {
        it('displays a create service button', () => {
            const wrapper = shallow(<ServiceList />)
            expect(wrapper.find('[data-test="create-service"]').exists()).toBe(
                true,
            )
        })

        it('should resolve fetchServices and set the result state accordingly', () => {
            const thePromise = Promise.resolve(mockServiceList)
            ServiceList.prototype.fetchServices = jest.fn(() => thePromise)

            const wrapper = shallow(<ServiceList />)
            return flushPromises().then(() => {
                expect(wrapper.state().result).toEqual(mockServiceList)
            })
        })
    })

    describe('when an error occured during fetching the services', () => {
        it('should display an error message', () => {
            const thePromise = Promise.reject(
                new Error('An arbitrary error occured.'),
            )
            ServiceList.prototype.fetchServices = jest.fn(() => thePromise)

            const wrapper = shallow(<ServiceList />)
            return flushPromises().then(() => {
                expect(wrapper.find('[data-test="error"]').exists()).toBe(true)
            })
        })
    })
})
