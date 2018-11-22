import React from 'react'
import { shallow } from 'enzyme'

import Spinner from './Spinner'

it('renders component with data-test-id', () => {
    let wrapper = shallow(<Spinner />)
    let el = wrapper.find('[data-test-id="app-loader"]')
    expect(el).toHaveLength(1)
})
