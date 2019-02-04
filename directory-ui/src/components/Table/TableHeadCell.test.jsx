import React from 'react'
import { shallow, mount } from 'enzyme'
import TableHeadCell from './TableHeadCell'

it('should match snapshot', () => {
  expect(shallow(<TableHeadCell/>)).toMatchSnapshot()
})

it('should render the contents', () => {
  const wrapper = mount(<TableHeadCell>Heading</TableHeadCell>)
  expect(wrapper.text()).toEqual('Heading')
})

describe('text alignment', () => {
  it('should use left alignment as default', () => {
    const wrapper = shallow(<tbody><TableHeadCell/></tbody>)
    expect(wrapper.find(TableHeadCell).prop('align')).toEqual('left')
  })

  it('should support center alignment', () => {
    const wrapper = shallow(<tbody><TableHeadCell align="center"/></tbody>)
    expect(wrapper.find(TableHeadCell).prop('align')).toEqual('center')
  })

  it('should support right alignment', () => {
    const wrapper = shallow(<tbody><TableHeadCell align="right"/></tbody>)
    expect(wrapper.find(TableHeadCell).prop('align')).toEqual('right')
  })
})
