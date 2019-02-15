import React from 'react'
import { mount, shallow } from "enzyme";
import TableBodyCell from './TableBodyCell'

const addToTable = tableBodyCell =>
  <table>{ tableBodyCell }</table>

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

it('should support left alignment', () => {
  const wrapper = addToTable(<TableBodyCell align="left" />)
  expect(shallow(wrapper)).toMatchSnapshot()
})

it('should support right alignment', () => {
  const wrapper = addToTable(<TableBodyCell align="right" />)
  expect(shallow(wrapper)).toMatchSnapshot()
})

it('should support center alignment', () => {
  const wrapper = addToTable(<TableBodyCell align="center" />)
  expect(shallow(wrapper)).toMatchSnapshot()
})
