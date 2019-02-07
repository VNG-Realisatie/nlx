import React from 'react'
import { mount } from 'enzyme'
import TableHeadCell from './TableHeadCell'

const insertTableHeadCellIntoValidTable = tableHeadCell =>
  <table>
    <tbody>
      <tr>
        { tableHeadCell }
      </tr>
    </tbody>
  </table>

it('should match snapshot', () => {
  const wrapper = mount(insertTableHeadCellIntoValidTable(<TableHeadCell/>))
  expect(wrapper).toMatchSnapshot()
})

it('should render the contents', () => {
  const wrapper = mount(insertTableHeadCellIntoValidTable(<TableHeadCell>Heading</TableHeadCell>))
  expect(wrapper.text()).toEqual('Heading')
})

describe('text alignment', () => {
  it('should use left alignment as default', () => {
    const wrapper = mount(insertTableHeadCellIntoValidTable(<TableHeadCell/>))
    expect(wrapper.find(TableHeadCell).prop('align')).toEqual('left')
  })

  it('should support center alignment', () => {
    const wrapper = mount(insertTableHeadCellIntoValidTable(<TableHeadCell align="center"/>))
    expect(wrapper.find(TableHeadCell).prop('align')).toEqual('center')
  })

  it('should support right alignment', () => {
    const wrapper = mount(insertTableHeadCellIntoValidTable(<TableHeadCell align="right"/>))
    expect(wrapper.find(TableHeadCell).prop('align')).toEqual('right')
  })
})
