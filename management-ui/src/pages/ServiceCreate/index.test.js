// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import { flushPromises } from '../../testHelpers'
import ServiceCreate from './index'

describe('ServiceCreate', () => {
    describe('when initially rendering the component', () => {
        it('should display a title', () => {
            const wrapper = shallow(<ServiceCreate />)
            expect(wrapper.find('h1').exists()).toBe(true)
        })
    })

    describe('when submitting the form', () => {
        it('should call the postService method onSubmit', () => {
            const thePromise = Promise.resolve(() => {})
            ServiceCreate.prototype.postService = jest.fn(() => thePromise)

            const wrapper = shallow(<ServiceCreate />)

            const instance = wrapper.instance()
            instance.onSubmit()

            return flushPromises().then(() => {
                expect(
                    ServiceCreate.prototype.postService,
                ).toHaveBeenCalledTimes(1)
            })
        })

        describe('the API call to save the service fails', () => {
            it('should display an error message when it occured', () => {
                const thePromise = Promise.reject(
                    new Error('An arbitrary error occured.'),
                )
                ServiceCreate.prototype.postService = jest.fn(() => thePromise)

                const wrapper = shallow(<ServiceCreate />)

                const instance = wrapper.instance()
                instance.onSubmit()

                return flushPromises().then(() => {
                    expect(wrapper.find('[data-test="error"]').exists()).toBe(
                        true,
                    )
                })
            })
        })
    })
})
