// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { Form, Formik, Field } from 'formik'
import { Select } from '@commonground/design-system'
import { func, string } from 'prop-types'
import { useHistory } from 'react-router-dom'
import {
  StyledFilters,
  StyledInput,
  StyledSearchIcon,
  StyledSwitch,
} from './index.styles'

const Filters = ({
  onQueryChanged,
  onStatusFilterChanged,
  queryValue,
  ...props
}) => {
  // TODO: Base default value of current environment
  // "development"

  const options = [
    {
      label: 'Acceptatie',
      value: 'https://directory.acc.nlx.io/',
    },
    {
      label: 'Demo',
      value: 'https://directory.demo.nlx.io/',
    },
    {
      label: 'Preproductie',
      value: 'https://directory.preprod.nlx.io/',
    },
    {
      label: 'Productie',
      value: 'https://directory.prod.nlx.io/',
    },
  ]

  const initialValue = {
    option: options[0].value,
  }

  const handleSelect = (e) => {
    window.location.href = e.value
  }

  return (
    // <StyledFilters {...props}>
    <Formik initialValues={initialValue} onSubmit={() => {}}>
      <Select
        options={options}
        name="option"
        placeholder="Selecteer omgeving"
        onQueryChanged={handleSelect}
        onChange={handleSelect}
      />
    </Formik>

    //   {/* <StyledSearchIcon />
    // <StyledInput
    //   value={queryValue}
    //   placeholder="Search for an organization or service…"
    //   onChange={(event) => onQueryChanged(event.target.value)}
    // />
    // <StyledSwitch
    //   id="include-offline-switch"
    //   label="Include offline"
    //   onChange={(event) => onStatusFilterChanged(event.target.checked)}
    // /> */}
    // </StyledFilters>
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
