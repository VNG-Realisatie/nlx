import React from 'react'
import { shallow } from 'enzyme'
import TableHead from './TableHead'

it('should match snapshot', () => {
  expect(shallow(<TableHead/>)).toMatchSnapshot()
})

it('should render child elements', () => {
  expect(shallow(<TableHead>
    <tr>
      <th>Table head</th>
    </tr>
  </TableHead>).contains(<tr>
    <th>Table head</th>
  </tr>)).toEqual(true)
})