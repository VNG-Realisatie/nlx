// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import ServiceList from './index'
import { Card } from './index.styles'

describe('ServiceList', () => {
    describe('when having no services', () => {
        let wrapper
        beforeAll(() => {
            wrapper = shallow(<ServiceList result={{ services: [] }} />)
        })

        it('displays a message', () => {
            expect(wrapper.text()).toBe('There are no services.')
        })
    })

    describe('when having two services', () => {
        let wrapper
        beforeAll(() => {
            wrapper = shallow(
                <ServiceList
                    result={{
                        services: [
                            { name: 'Service 1' },
                            { name: 'Service 2' },
                        ],
                    }}
                />,
            )
        })

        it('shows two Cards', () => {
            expect(wrapper.find(Card)).toHaveLength(2)
        })
    })
})
