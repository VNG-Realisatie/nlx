import React from 'react'
import Enzyme, { shallow } from 'enzyme'
import EnzymeAdapter from 'enzyme-adapter-react-16'

import Mock from './Mock'

Enzyme.configure({ adapter: new EnzymeAdapter() })

let comp

beforeEach(() => {
    comp = shallow(<Mock />)
})

it('renders component with data-test-id without error', () => {
    let el = comp.find('[data-test-id="mock-component"]')
    expect(el).toHaveLength(1)
})

it('has initial counter with value 0', () => {
    let cnt = comp.state('counter')
    expect(cnt).toBe(0)
})
