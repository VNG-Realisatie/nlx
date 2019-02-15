import React from 'react'
import { shallow } from 'enzyme'
import BodyCell from './BodyCell'

const addToTable = tableBodyCell =>
  <table>{ tableBodyCell }</table>

it('should match snapshot', () => {
  expect(shallow(<BodyCell/>)).toMatchSnapshot()
})

it('should render child elements', () => {
  expect(shallow(<BodyCell>
    <tr>
      <td>Table body</td>
    </tr>
  </BodyCell>).contains(<tr>
    <td>Table body</td>
  </tr>)).toEqual(true)
})

it('should support left alignment', () => {
  const wrapper = addToTable(<BodyCell align="left" />)
  expect(shallow(wrapper)).toMatchSnapshot()
})

it('should support right alignment', () => {
  const wrapper = addToTable(<BodyCell align="right" />)
  expect(shallow(wrapper)).toMatchSnapshot()
})

it('should support center alignment', () => {
  const wrapper = addToTable(<BodyCell align="center" />)
  expect(shallow(wrapper)).toMatchSnapshot()
})
