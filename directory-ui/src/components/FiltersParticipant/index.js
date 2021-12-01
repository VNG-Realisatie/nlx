// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//
import React from 'react'
import { SelectComponent } from '@commonground/design-system'
import { func, string } from 'prop-types'
import { ReactComponent as SearchIcon } from '../../icons/search.svg'
import { StyledFilters, StyledInput } from './index.styles'

const FiltersParticipant = ({ onQueryChanged, queryValue, ...props }) => {
  const options = [
    {
      label: 'Demo',
      value: 'https://directory.demo.nlx.io/participants',
    },
    {
      label: 'Pre-productie',
      value: 'https://directory.preprod.nlx.io/participants',
    },
    {
      label: 'Productie',
      value: 'https://directory.prod.nlx.io/participants',
    },
  ]

  const currentEnvironmentOption = () => {
    const environment = window.location.hostname

    // eslint-disable-next-line default-case
    switch (true) {
      case /demo/.test(environment):
        return options[0]
      case /preprod/.test(environment):
        return options[1]
      case /prod/.test(environment):
        return options[2]
      case /directory.nlx.io/.test(environment):
        return options[0]
    }
  }

  const handleSelect = (e) => {
    window.location.href = e.value
  }

  return (
    <StyledFilters {...props}>
      <SelectComponent
        options={options}
        defaultValue={currentEnvironmentOption()}
        name="option"
        placeholder="Selecteer omgeving"
        onQueryChanged={handleSelect}
        onChange={handleSelect}
      />
      <StyledInput
        type="text"
        name="search"
        placeholder="Zoeken…"
        icon={SearchIcon}
        onChange={(event) => onQueryChanged(event.target.value)}
        defaultValue={queryValue}
      />
    </StyledFilters>
  )
}

FiltersParticipant.propTypes = {
  onQueryChanged: func,
  queryValue: string,
}

FiltersParticipant.defaultProps = {
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  onQueryChanged: () => {},
}

export default FiltersParticipant
