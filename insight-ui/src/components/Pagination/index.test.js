import React from 'react'
import ReactDOM from 'react-dom'
import {mount} from 'enzyme'
import Pagination, { hasNextPage, hasPreviousPage, calcAmountOfPages } from './index'
import { StyledButton } from './index.styles'

it('renders without crashing', () => {
  expect(() => {
    const div = document.createElement('div')
    ReactDOM.render(<Pagination currentPage={1} totalRows={20} rowsPerPage={10} />, div)
  }).not.toThrow()
})

test('has previous page', () => {
  expect(hasPreviousPage(1)).toEqual(false)
  expect(hasPreviousPage(2)).toEqual(true)
})

test('has next page', () => {
  expect(hasNextPage(1, 42)).toEqual(true)
  expect(hasNextPage(42, 42)).toEqual(false)
})

test.each(
  [[1, 1, 1], [20, 10, 2], [10, 0, 0], [0, 10, 0]]
)('calculate amount of pages for %s rows and %s rows per page', (totalRows, rowsPerPage, expected) => {
  expect(calcAmountOfPages(totalRows, rowsPerPage)).toBe(expected)
})

describe('changing the input value', () => {
  it('should trigger the onPageChangedHandler function', () => {
    const onPageChangedHandler = jest.fn()
    const wrapper = mount(<Pagination currentPage={1} totalRows={20} rowsPerPage={10} onPageChangedHandler={onPageChangedHandler} />)
    const input = wrapper.find('input')
    input.simulate('change', { target: { value: 2 } })
    expect(onPageChangedHandler).toHaveBeenCalledWith(2)
  })
})

describe('clicking the next page button', () => {
  it('should trigger the onPageChanged callback', () => {
    const onPageChangedHandler = jest.fn()
    const wrapper = mount(<Pagination currentPage={1} totalRows={20} rowsPerPage={10} onPageChangedHandler={onPageChangedHandler} />)
    const nextButton = wrapper.find(StyledButton).last()
    nextButton.simulate('click')
    expect(onPageChangedHandler).toHaveBeenCalledWith(2)
  })
})
