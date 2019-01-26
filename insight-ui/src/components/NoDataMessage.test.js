import React from 'react'
import { mount } from 'enzyme'

import NoDataMessage from './NoDataMessage'

let component, msg

beforeAll(() => {
    component = mount(<NoDataMessage />)
})

beforeEach(() => {
    msg = component.text('[data-test-id="no-data-msg"]')
})

describe('<NoDataMessage />', () => {
    it('mounts NoDataMessage component', () => {
        expect(component).toBeTruthy()
    })

    it('shows default text with length > 3', () => {
        expect(msg.length).toBeGreaterThan(3)
    })

    it('shows prop.msg received', () => {
        const testProp = { msg: 'Test messsage!' }
        component.setProps(testProp)
        msg = component.text('[data-test-id="no-data-msg"]')
        expect(msg).toContain(testProp.msg)
    })
})
