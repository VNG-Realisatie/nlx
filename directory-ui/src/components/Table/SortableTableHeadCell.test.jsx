import React from 'react'
import { shallow, mount } from 'enzyme'
import SortableTableHeadCell from './SortableTableHeadCell'

const addToTable = headCell =>
  <table><thead><tr>{headCell}</tr></thead></table>

it('should match snapshot', () => {
  expect(shallow(<SortableTableHeadCell/>)).toMatchSnapshot()
})

it('should support the ascending direction', () => {
  const wrapper = addToTable(<SortableTableHeadCell direction="asc"/>)
  expect(mount(wrapper)).toMatchSnapshot()
})

it('should support the descending direction', () => {
  const wrapper = addToTable(<SortableTableHeadCell direction="desc"/>)
  expect(mount(wrapper)).toMatchSnapshot()
})
