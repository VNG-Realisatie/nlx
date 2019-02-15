import React from 'react'
import { shallow } from 'enzyme'
import Head from './Head'

it('should match snapshot', () => {
  expect(shallow(<Head/>)).toMatchSnapshot()
})

it('should render child elements', () => {
  expect(shallow(<Head>
    <tr>
      <th>Table head</th>
    </tr>
  </Head>).contains(<tr>
    <th>Table head</th>
  </tr>)).toEqual(true)
})