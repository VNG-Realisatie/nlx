// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import { Formik } from 'formik'
import ServiceForm, { validationSchema } from './index'
import { ValidationError } from 'yup'

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

describe('the form validator', () => {
    it('should be valid with a NON fully qualified domain name as endpointURL and apiSpecificationURL', () => {
        var result = validationSchema.validate({
            name: 'service.name',
            endpointURL: 'http://service:8080',
            apiSpecificationURL: 'http://service:8080/openapispec.yml',
        })
        return expect(result).resolves.toBeDefined()
    })

    it('should be valid with an IP address as endpointURL and apiSpecificationURL', () => {
        var result = validationSchema.validate({
            name: 'service.name',
            endpointURL: 'http://10.0.0.1:8080',
            apiSpecificationURL: 'http://10.0.0.1:8080/openapispec.yml',
        })
        return expect(result).resolves.toBeDefined()
    })

    it('should not be valid if the endpointURL is missing', () => {
        var result = validationSchema.validate({
            name: 'service.name',
            endpointURL: '',
            apiSpecificationURL: 'http://service:8080/openapispec.yml',
        })
        return expect(result).rejects.toEqual(
            new ValidationError('endpointURL is a required field'),
        )
    })
})
