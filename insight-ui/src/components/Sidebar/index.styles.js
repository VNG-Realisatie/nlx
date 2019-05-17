import styled from "styled-components";
import { Search } from "@commonground/design-system";

export const StyledSidebar = styled.div`
  flex: 0 0 200px;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45,50,64,.05), 0 1px 8px rgba(45,50,64,.05);
  z-index: 1;
  min-height: calc(100% - 56px);
`

export const StyledSearch = styled(Search)`
  box-shadow: none;
  border-bottom: 1px solid #F0F2F7;

  input {
    font-size: 14px;
  }
`

export const StyledOrganizationList = styled.ul`
  list-style-type: none;
  padding: 0;
  
  a {
    color: #A3AABF;
    font-weight: 600;
    text-decoration: none;
    display: block;
    padding: 5px 0 5px 20px;
    line-height: 22px;
    font-size: 14px;

    
    &.active {
      color: #517FFF;
      border-left: 2px solid #517FFF;
      padding-left: 18px;
    }
  }
`
