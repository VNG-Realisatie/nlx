// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import { flushPromises } from '../../testHelpers'
import InwayList from './index'

const mockInwayList = {
    inways: [{ name: 'inway1' }, { name: 'inway2' }],
}

describe('InwayList', () => {
    describe('the component is initialized', () => {
        it('should resolve fetchInways and set the result state accordingly', () => {
            const thePromise = Promise.resolve(mockInwayList)
            InwayList.prototype.fetchInways = jest.fn(() => thePromise)

            const wrapper = shallow(<InwayList />)
            return flushPromises().then(() => {
                expect(wrapper.state().result).toEqual(mockInwayList)
            })
        })
    })

    describe('when an error occured during fetching the apis', () => {
        it('should show an error message', () => {
            const thePromise = Promise.reject(
                new Error('An arbitrary error occured.'),
            )
            InwayList.prototype.fetchInways = jest.fn(() => thePromise)

            const wrapper = shallow(<InwayList />)
            return flushPromises().then(() => {
                expect(wrapper.find('[data-test="error"]').exists()).toBe(true)
            })
        })
    })
})
