import React from 'react'
import { shallow } from 'enzyme'

import NoDataMessage from './NoDataMessage'

describe('the <NoDataMessage /> component', () => {
    it('mounts successfully', () => {
        const component = shallow(<NoDataMessage />)
        expect(component).toBeTruthy()
    })

    describe('when no message is provided', () => {
        it('should render a default message', () => {
            const component = shallow(<NoDataMessage />)
            const message = component.find('[data-test="message"]')
            expect(message.text()).toBe('No logs to show')
        })
    })

    describe('when a message is provided', () => {
        it('should show the provided message', () => {
            const component = shallow(<NoDataMessage msg="Test message!" />)
            const message = component.find('[data-test="message"]')
            expect(message.text()).toBe('Test message!')
        })
    })
})
