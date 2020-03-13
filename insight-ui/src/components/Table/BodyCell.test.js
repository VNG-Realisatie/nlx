import React from 'react'
import BodyCell from './BodyCell'

xtest('should render child elements', () => {
  expect(
    shallow(
      <BodyCell>
        <tr>
          <td>Table body</td>
        </tr>
      </BodyCell>,
    ).contains(
        <td>Table body</td>
                            </tr>)).toEqual(true)
})
