import styled from 'styled-components'
import SearchIcon from './SearchIcon';
import Switch from '../Switch/Switch';

export const StyledFilters = styled.div`
  background: #ffffff;
  border: 1px solid #EAECF0;
  border-radius: 3px;
  display: flex;
  align-items: center;
  padding: 0 1.5rem;
`

export const StyledSearchIcon = styled(SearchIcon)`
  display: inline-block;
  width: 12px;
  height: 12px;
  margin-right: 1rem;
`

export const StyledInput = styled.input`
  border: 0 none;
  flex-grow: 1;
  font-size: 1rem;
  font-family: 'Source Sans Pro', sans-serif;
  height: 50px;
  
  &::placeholder {
    color: #CAD0E0;
  }
`

export const StyledSwitch = styled(Switch)`
  vertical-align: middle;
`
