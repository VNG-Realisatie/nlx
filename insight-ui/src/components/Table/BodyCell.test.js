import React from 'react'
import { shallow } from 'enzyme'
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
      <tr>
        <td>Table body</td>
      </tr>,
    ),
  ).toEqual(true)
})
