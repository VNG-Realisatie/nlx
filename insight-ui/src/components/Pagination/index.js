// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, number } from 'prop-types'
import CaretLeft from './CaretLeft'
import CaretRight from './CaretRight'
import {
  StyledPagination,
  StyledButton,
  StyledInput,
  StyledLabel,
} from './index.styles'

const noop = () => {}

export const hasPreviousPage = (currentPage) => currentPage > 1

export const hasNextPage = (currentPage, amountOfPages) =>
  currentPage < amountOfPages

export const calcAmountOfPages = (totalRows, rowsPerPage) =>
  rowsPerPage !== 0 ? Math.ceil(totalRows / rowsPerPage) : 0

export const onPreviousPageButtonClickedHandler = (
  currentPage,
  onPageChangedHandler,
) =>
  hasPreviousPage(currentPage) ? onPageChangedHandler(currentPage - 1) : noop

export const onNextPageButtonClickedHandler = (
  currentPage,
  amountOfPages,
  onPageChangedHandler,
) =>
  hasNextPage(currentPage, amountOfPages)
    ? onPageChangedHandler(currentPage + 1)
    : noop

const Pagination = ({
  currentPage,
  totalRows,
  rowsPerPage,
  onPageChangedHandler,
  ...props
}) => {
  const amountOfPages = calcAmountOfPages(totalRows, rowsPerPage)

  const handlePreviousPageButtonClicked = () =>
    onPreviousPageButtonClickedHandler(currentPage, onPageChangedHandler)

  const handleNextPageButtonClicked = () =>
    onNextPageButtonClickedHandler(
      currentPage,
      amountOfPages,
      onPageChangedHandler,
    )
  return (
    <StyledPagination {...props}>
      <StyledButton
        onClick={handlePreviousPageButtonClicked}
        disabled={!hasPreviousPage(currentPage)}
        aria-label="Vorige pagina"
      >
        <CaretLeft
          color={hasPreviousPage(currentPage) ? '#2D3240' : '#CAD0E0'}
          focusable={false}
        />
      </StyledButton>
      <StyledLabel>
        Pagina{' '}
        <StyledInput
          type="number"
          value={currentPage}
          min={1}
          max={amountOfPages}
          onChange={(event) => onPageChangedHandler(event.target.value)}
        />
        van {amountOfPages}
      </StyledLabel>
      <StyledButton
        onClick={handleNextPageButtonClicked}
        disabled={!hasNextPage(currentPage, amountOfPages)}
        aria-label="Volgende pagina"
      >
        <CaretRight
          color={
            hasNextPage(currentPage, amountOfPages) ? '#2D3240' : '#CAD0E0'
          }
          focusable={false}
        />
      </StyledButton>
    </StyledPagination>
  )
}

Pagination.propTypes = {
  onPageChangedHandler: func,
  currentPage: number,
  totalRows: number.isRequired,
  rowsPerPage: number.isRequired,
}

Pagination.defaultProps = {
  onPageChangedHandler: noop,
}

export default Pagination
