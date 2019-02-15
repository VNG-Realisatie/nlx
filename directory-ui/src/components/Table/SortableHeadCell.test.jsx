import React from 'react'
import { shallow, mount } from 'enzyme'
import SortableHeadCell from './SortableHeadCell'

const addToTable = headCell =>
  <table><thead><tr>{headCell}</tr></thead></table>

it('should match snapshot', () => {
  expect(shallow(<SortableHeadCell/>)).toMatchSnapshot()
})

it('should support the ascending direction', () => {
  const wrapper = addToTable(<SortableHeadCell direction="asc"/>)
  expect(mount(wrapper)).toMatchSnapshot()
})

it('should support the descending direction', () => {
  const wrapper = addToTable(<SortableHeadCell direction="desc"/>)
  expect(mount(wrapper)).toMatchSnapshot()
})
