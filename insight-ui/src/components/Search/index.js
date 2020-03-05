import React from 'react'
import { object, func } from 'prop-types'
import { StyledSearch, StyledInput, StyledSearchIcon } from './index.styles'

const Search = ({ inputProps, onQueryChanged, children, ...props }) =>
  <StyledSearch {...props}>
    <StyledSearchIcon />
    <StyledInput dataTest="query"
                 onChange={event => onQueryChanged(event.target.value)}
                 {...inputProps}
    />
    {children}
  </StyledSearch>

Search.propTypes = {
  onQueryChanged: func,
  inputProps: object
}

Search.defaultProps = {
  onQueryChanged: () => {},
  inputProps: {
    placeholder: 'Search…'
  }
}

export default Search
