import React from 'react'
import { oneOf } from 'prop-types'
import { StyledSortableTableHeadCell, StyledArrow } from './SortableHeadCell.styles'

export const ASCENDING = 'asc'
export const DESCENDING = 'desc'

const ArrowUp = () =>
  <StyledArrow viewBox="0 0 8 10">
    <path d="M5 6V0H3v6H0l4 4 4-4z" fill="#2D3240" />
  </StyledArrow>

const ArrowDown = () =>
  <StyledArrow viewBox="0 0 8 10">
    <path d="M3 4v6h2V4h3L4 0 0 4z" fill="#2D3240" />
  </StyledArrow>

const SortableArrow = ({ direction }) =>
  direction === ASCENDING ? <ArrowDown/> : <ArrowUp/>

const SortableHeadCell = ({ children, direction, ...other }) =>
  <StyledSortableTableHeadCell isSorting={direction} {...other}>
    {children}
    {direction ? <SortableArrow direction={direction} /> : null}
  </StyledSortableTableHeadCell>

SortableHeadCell.propTypes = {
  direction: oneOf([ASCENDING, DESCENDING, null])
}

export default SortableHeadCell
