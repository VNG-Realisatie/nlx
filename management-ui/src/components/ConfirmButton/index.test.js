// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'
import { shallow } from 'enzyme'
import ConfirmButton from './index'

describe('ConfirmButton', () => {
    let callbackFn
    let wrapper

    beforeEach(() => {
        callbackFn = jest.fn()
        wrapper = shallow(<ConfirmButton onConfirm={callbackFn} />)
    })

    it('should exist', () => {
        expect(wrapper.exists()).toBe(true)
    })

    it('should have a delete button', () => {
        const button = wrapper.find('[data-test="delete-button"]')
        expect(button.exists()).toBe(true)
    })

    describe('showing the confirmation', () => {
        beforeEach(() => {
            const deleteButton = wrapper.find('[data-test="delete-button"]')
            deleteButton.simulate('click')
        })

        it('should show a confirm and cancel button', () => {
            const cancelButton = wrapper.find('[data-test="cancel-button"]')
            expect(cancelButton.exists()).toBe(true)

            const confirmButton = wrapper.find('[data-test="confirm-button"]')
            expect(confirmButton.exists()).toBe(true)
        })

        describe('confirming', () => {
            it('calls the callback', () => {
                const confirmButton = wrapper.find(
                    '[data-test="confirm-button"]',
                )
                confirmButton.simulate('click')

                expect(callbackFn).toHaveBeenCalledTimes(1)
            })
        })

        describe('cancelling', () => {
            beforeEach(() => {
                const cancelButton = wrapper.find('[data-test="cancel-button"]')
                cancelButton.simulate('click')
            })

            it('toggles back to the initial state', () => {
                const deleteButton = wrapper.find('[data-test="delete-button"]')
                expect(deleteButton.exists()).toBe(true)
            })
        })
    })
})
