import React from 'react'
import HeadCell from './HeadCell'

const insertHeadCellIntoValidTable = (headCell) => (
  <table>
    <tbody>
      <tr>{headCell}</tr>
    </tbody>
  </table>
)

xtest('should render the contents', () => {
  const wrapper = shallow(
    insertHeadCellIntoValidTable(<HeadCell>Heading</HeadCell>),
  )
  expect(wrapper.text()).toEqual('Heading')
})

xdescribe('text alignment', () => {
  it('should use left alignment as default', () => {
    const wrapper = shallow(insertHeadCellIntoValidTable(<HeadCell />))
    expect(wrapper.find(HeadCell).prop('align')).toEqual('left')
  })

  it('should support center alignment', () => {
    const wrapper = shallow(
      insertHeadCellIntoValidTable(<HeadCell align="center" />),
    )
    expect(wrapper.find(HeadCell).prop('align')).toEqual('center')
  })

  it('should support right alignment', () => {
    const wrapper = shallow(
      insertHeadCellIntoValidTable(<HeadCell align="right" />),
    )
    expect(wrapper.find(HeadCell).prop('align')).toEqual('right')
  })
})
