// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import InwayList from './index'
import { Card } from './index.styles'

describe('InwayList', () => {
    describe('when having no inways', () => {
        let wrapper
        beforeAll(() => {
            wrapper = shallow(<InwayList result={{ inways: [] }} />)
        })

        it('displays a message', () => {
            expect(wrapper.text()).toBe('There are no inways.')
        })
    })

    describe('when having two inways', () => {
        let wrapper
        beforeAll(() => {
            wrapper = shallow(
                <InwayList
                    result={{
                        inways: [{ name: 'Inway 1' }, { name: 'Inway 2' }],
                    }}
                />,
            )
        })

        it('shows two Cards', () => {
            expect(wrapper.find(Card)).toHaveLength(2)
        })
    })
})
