import React from 'react'
import { shallow } from 'enzyme'
import Table from './index'

it('should match snapshot', () => {
  expect(shallow(<Table/>)).toMatchSnapshot()
})

it('should render child elements', () => {
  expect(shallow(<Table>
    <tr>
      <td>Table body</td>
    </tr>
  </Table>).contains(<tr>
    <td>Table body</td>
  </tr>)).toEqual(true)
})
