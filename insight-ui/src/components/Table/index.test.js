// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shallow } from 'enzyme'
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
