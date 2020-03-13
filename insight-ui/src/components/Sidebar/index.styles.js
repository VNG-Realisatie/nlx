// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import styled from 'styled-components'
import Search from '../Search'

export const StyledSidebar = styled.div`
  flex: 0 0 200px;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45, 50, 64, 0.05), 0 1px 8px rgba(45, 50, 64, 0.05);
  z-index: 1;
  min-height: calc(100% - 56px);
`

export const StyledSearch = styled(Search)`
  box-shadow: none;
  border-bottom: 1px solid #f0f2f7;

  input {
    font-size: 14px;
  }
`

export const StyledOrganizationList = styled.ul`
  list-style-type: none;
  padding: 0;

  a {
    color: #a3aabf;
    font-weight: 600;
    text-decoration: none;
    display: block;
    padding: 5px 0 5px 20px;
    line-height: 22px;
    font-size: 14px;

    &:hover {
      background-color: #f7f9fc;
    }

    &:active {
      background-color: #f0f2f7;
    }

    &.active {
      color: #517fff;
      border-left: 2px solid #517fff;
      padding-left: 18px;
    }
  }
`
