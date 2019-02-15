import React from 'react'
import { mount, shallow } from 'enzyme'
import HeadCell from './HeadCell'

const insertHeadCellIntoValidTable = headCell =>
  <table>
    <tbody>
      <tr>
        { headCell }
      </tr>
    </tbody>
  </table>

it('should match snapshot', () => {
  const wrapper = mount(insertHeadCellIntoValidTable(<HeadCell/>))
  expect(wrapper).toMatchSnapshot()
})

it('should render the contents', () => {
  const wrapper = shallow(insertHeadCellIntoValidTable(<HeadCell>Heading</HeadCell>))
  expect(wrapper.text()).toEqual('Heading')
})

describe('text alignment', () => {
  it('should use left alignment as default', () => {
    const wrapper = shallow(insertHeadCellIntoValidTable(<HeadCell/>))
    expect(wrapper.find(HeadCell).prop('align')).toEqual('left')
  })

  it('should support center alignment', () => {
    const wrapper = shallow(insertHeadCellIntoValidTable(<HeadCell align="center"/>))
    expect(wrapper.find(HeadCell).prop('align')).toEqual('center')
  })

  it('should support right alignment', () => {
    const wrapper = shallow(insertHeadCellIntoValidTable(<HeadCell align="right"/>))
    expect(wrapper.find(HeadCell).prop('align')).toEqual('right')
  })
})
