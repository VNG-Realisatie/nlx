// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import { flushPromises } from '../../testHelpers'
import ServiceUpdate from './index'

const mockService = {
    name: 'service1',
}

const mockMatch = {
    params: {
        name: 'service1',
    },
}

describe('ServiceUpdate', () => {
    describe('the component is initialized', () => {
        it('should resolve fetchService and set the result state accordingly', () => {
            const thePromise = Promise.resolve(mockService)
            ServiceUpdate.prototype.fetchService = jest.fn(() => thePromise)

            const wrapper = shallow(<ServiceUpdate match={mockMatch} />)
            return flushPromises().then(() => {
                expect(wrapper.state().data).toEqual(mockService)
            })
        })

        describe('when an error occured', () => {
            it('displays an error', () => {
                const thePromise = Promise.reject(
                    new Error('An arbitrary error occured.'),
                )
                ServiceUpdate.prototype.fetchService = jest.fn(() => thePromise)

                const wrapper = shallow(<ServiceUpdate match={mockMatch} />)
                return flushPromises().then(() => {
                    expect(wrapper.find('[data-test="error"]').exists()).toBe(
                        true,
                    )
                })
            })
        })
    })

    describe('the onSubmit is called', () => {
        it('calls putService', () => {
            ServiceUpdate.prototype.putService = jest.fn(() =>
                Promise.resolve(mockService),
            )

            const wrapper = shallow(<ServiceUpdate match={mockMatch} />)
            wrapper.setState({ data: { name: 'service1' } })

            const instance = wrapper.instance()
            instance.onSubmit(mockService)

            return flushPromises().then(() => {
                expect(ServiceUpdate.prototype.putService).toHaveBeenCalledWith(
                    'service1',
                    mockService,
                )
            })
        })
    })

    describe('the onDelete is called', () => {
        it('calls deleteService', () => {
            ServiceUpdate.prototype.deleteService = jest.fn(() =>
                Promise.resolve({}),
            )

            const wrapper = shallow(<ServiceUpdate match={mockMatch} />)
            wrapper.setState({ data: { name: 'service1' } })

            const instance = wrapper.instance()
            instance.onDelete(mockService)

            return flushPromises().then(() => {
                expect(
                    ServiceUpdate.prototype.deleteService,
                ).toHaveBeenCalledWith('service1')
            })
        })
    })
})
