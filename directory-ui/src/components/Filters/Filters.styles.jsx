import styled from 'styled-components'
import Card from '../Card/Card';
import SearchIcon from './SearchIcon';
import Switch from '../Switch/Switch';

export const StyledFilters = styled(Card)`
  display: flex;
  align-items: center;
  padding: 0 24px 1px 20px;
`

export const StyledSearchIcon = styled(SearchIcon)`
  flex-shrink: 0;
  width: 12px;
  height: 12px;
  margin-top: 1px;
  margin-right: 12px;
`

export const StyledInput = styled.input`
  flex-grow: 1;
  border: none;
  padding: 0;
  font-size: 1rem;
  font-family: 'Source Sans Pro', sans-serif;
  height: 48px;

  &::placeholder {
    color: #CAD0E0;
  }
`

export const StyledSwitch = styled(Switch)`
  flex-shrink: 0;
`