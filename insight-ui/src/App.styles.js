import styled from 'styled-components'
import Card from './components/Card'
import { Navbar, Search } from '@commonground/design-system'

export default styled.div`
  display: flex;
  height: 100vh;
  flex-wrap: wrap;
  align-content: flex-start;
`

export const StyledNavbar = styled(Navbar)`
  flex: 1 100%;
  z-index: 3;
`

export const StyledSidebar = styled.div`
  flex: 0 0 200px;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45,50,64,.05), 0 1px 8px rgba(45,50,64,.05);
  z-index: 1;
  height: calc(100% - 56px);
`

export const StyledSearch = styled(Search)`
  box-shadow: none;
  border-bottom: 1px solid #F0F2F7;
`

export const StyledContent = styled.div`
  flex: 1;
  background: #F7F9FC;
  display: flex;
  justify-content: center;
  align-items: center;
`

export const StyledOrganizationList = styled.ul`
  list-style-type: none;
  padding: 0;
  
  li {
    &.active {
      color: #517FFF;
    
      a {
        border-left: 2px solid #517FFF;
      }
    }
  }
  
  a {
    color: #A3AABF;
    font-weight: 600;
    text-decoration: none;
    display: block;
    padding-left: 20px;
    line-height: 32px;
  }
`

export const StyledCard = styled(Card)`
  width: 400px;
  padding: 8px 24px 8px 24px;
  
  .text-muted {
    font-size: 12px;
    color: #A3AABF;
    
    a {
      text-decoration: none;
      color: #517FFF;
    }
  }
`
