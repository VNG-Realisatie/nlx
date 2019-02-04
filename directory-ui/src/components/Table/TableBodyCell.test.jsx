import React from 'react'
import { shallow } from 'enzyme'
import TableBodyCell from './TableBodyCell'

it('should match snapshot', () => {
  expect(shallow(<TableBodyCell/>)).toMatchSnapshot()
})

it('should render child elements', () => {
  expect(shallow(<TableBodyCell>
    <tr>
      <td>Table body</td>
    </tr>
  </TableBodyCell>).contains(<tr>
    <td>Table body</td>
  </tr>)).toEqual(true)
})
