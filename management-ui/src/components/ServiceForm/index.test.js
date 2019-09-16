// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import { Formik } from 'formik'
import ServiceForm from './index'

describe('ServiceForm', () => {
    let onSubmit, wrapper
    beforeAll(() => {
        onSubmit = jest.fn()
        wrapper = shallow(<ServiceForm onSubmit={onSubmit} />)
    })

    it('should exist', () => {
        expect(wrapper.exists()).toBe(true)
    })

    describe('form submitted successfully', () => {
        it('should call the onSubmit function', () => {
            const form = wrapper.find(Formik)
            form.simulate('submit')

            expect(onSubmit).toHaveBeenCalledTimes(1)
        })
    })
})
