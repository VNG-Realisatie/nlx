import React from 'react'
import Table from './index'

xtest('should render child elements', () => {
  expect(
    shallow(
      <Table>
        <tr>
          <td>Table body</td>
        </tr>
      </Table>,
    ).contains(
      <tr>
        <td>Table body</td>
      </tr>,
    ),
  ).toEqual(true)
})
