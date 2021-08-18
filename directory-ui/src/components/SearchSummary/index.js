// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React from 'react'
import { number } from 'prop-types'
import { Text } from './index.styles'

const SearchSummary = (props) => {
  const { totalServices, totalFilteredServices } = props

  return (
    <Text>
      {`${
        totalFilteredServices !== totalServices
          ? `${totalFilteredServices} van `
          : ''
      }
      ${totalServices} beschikbare service${
        totalServices > 1 ? 's' : ''
      }`.toUpperCase()}
    </Text>
  )
}

SearchSummary.propTypes = {
  totalServices: number.isRequired,
  totalFilteredServices: number.isRequired,
}

export default SearchSummary
