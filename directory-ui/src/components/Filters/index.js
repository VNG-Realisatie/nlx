// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//
import React from 'react'
import { SelectComponent } from '@commonground/design-system'
import { func, string } from 'prop-types'
import { ReactComponent as SearchIcon } from '../../icons/search.svg'
import Switch from '../Switch'
import { StyledFilters, StyledInput } from './index.styles'

const Filters = ({
  onQueryChanged,
  onStatusFilterChanged,
  queryValue,
  ...props
}) => {
  const options = [
    {
      label: 'Demo',
      value: 'https://directory.demo.nlx.io/',
    },
    {
      label: 'Pre-productie',
      value: 'https://directory.preprod.nlx.io/',
    },
    {
      label: 'Productie',
      value: 'https://directory.prod.nlx.io/',
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
      <Switch
        id="include-offline-switch"
        label="Toon offline services"
        onChange={(event) => onStatusFilterChanged(event.target.checked)}
      />
    </StyledFilters>
  )
}

Filters.propTypes = {
  onQueryChanged: func,
  onStatusFilterChanged: func,
  queryValue: string,
}

Filters.defaultProps = {
  onQueryChanged: () => {},
  onStatusFilterChanged: () => {},
}

export default Filters
