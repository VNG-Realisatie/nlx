import React from 'react'
import BodyCell from './BodyCell'

xit('should render child elements', () => {
  expect(shallow(<BodyCell>
    <tr>
      <td>Table body</td>
    </tr>
  </BodyCell>).contains(<tr>
    <td>Table body</td>
  </tr>)).toEqual(true)
});
