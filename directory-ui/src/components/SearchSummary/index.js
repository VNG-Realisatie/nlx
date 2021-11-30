// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { number, string } from 'prop-types'
import { Text } from './index.styles'

const SearchSummary = (props) => {
  const {
    totalItems,
    totalFilteredItems,
    itemDescription,
    itemPluralDescription,
  } = props

  return (
    <Text>
      {`${totalFilteredItems !== totalItems ? `${totalFilteredItems} van ` : ''}
      ${totalItems} ${
        totalItems > 1 ? itemPluralDescription : itemDescription
      }`.toUpperCase()}
    </Text>
  )
}

SearchSummary.propTypes = {
  totalItems: number.isRequired,
  totalFilteredItems: number.isRequired,
  itemDescription: string.isRequired,
  itemPluralDescription: string.isRequired,
}

export default SearchSummary
